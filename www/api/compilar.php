<?php
/**
 * Verbo Portal — API de Compilação e Execução
 *
 * Recebe código Verbo via POST JSON, executa usando o binário `verbo`,
 * e retorna o output como JSON.
 *
 * Integração Yii2: Adaptar para um Controller/Action.
 *
 * Exemplo de request:
 *   POST /api/compilar.php
 *   Content-Type: application/json
 *   {"codigo": "Exibir com (\"Olá!\")."}
 *
 * Exemplo de response:
 *   {"output": "Olá!\n", "erro": "", "tempo": 0.35}
 */

header('Content-Type: application/json; charset=utf-8');
header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: POST, OPTIONS');
header('Access-Control-Allow-Headers: Content-Type');

// Handle CORS preflight
if ($_SERVER['REQUEST_METHOD'] === 'OPTIONS') {
    http_response_code(204);
    exit;
}

if ($_SERVER['REQUEST_METHOD'] !== 'POST') {
    http_response_code(405);
    echo json_encode(['erro' => 'Método não permitido. Use POST.']);
    exit;
}

// Parse input
$input = json_decode(file_get_contents('php://input'), true);
if (!$input || empty($input['codigo'])) {
    http_response_code(400);
    echo json_encode(['erro' => 'Campo "codigo" é obrigatório.']);
    exit;
}

$codigo = $input['codigo'];

// Validação básica de segurança
if (strlen($codigo) > 10000) {
    http_response_code(400);
    echo json_encode(['erro' => 'Código muito longo. Máximo 10.000 caracteres.']);
    exit;
}

// Caminho do binário Verbo — ajustar conforme ambiente
$verboBin = realpath(__DIR__ . '/../../build/verbo');
if (!$verboBin || !is_executable($verboBin)) {
    // Fallback: tentar no PATH
    $verboBin = 'verbo';
}

// Criar arquivo temporário .vrb
$tmpDir = sys_get_temp_dir();
$tmpFile = tempnam($tmpDir, 'verbo_') . '.vrb';
file_put_contents($tmpFile, $codigo);

// Executar com timeout
$startTime = microtime(true);
$timeout = 10; // segundos

$descriptors = [
    0 => ['pipe', 'r'],  // stdin
    1 => ['pipe', 'w'],  // stdout
    2 => ['pipe', 'w'],  // stderr
];

$process = proc_open(
    escapeshellcmd($verboBin) . ' executar ' . escapeshellarg($tmpFile),
    $descriptors,
    $pipes,
    $tmpDir,
    ['PATH' => getenv('PATH'), 'HOME' => getenv('HOME'), 'GOPATH' => getenv('GOPATH')]
);

$output = '';
$erro = '';

if (is_resource($process)) {
    fclose($pipes[0]); // Fechar stdin

    // Ler com timeout
    stream_set_blocking($pipes[1], false);
    stream_set_blocking($pipes[2], false);

    $deadline = $startTime + $timeout;

    while (true) {
        $read = [$pipes[1], $pipes[2]];
        $write = null;
        $except = null;

        $remaining = $deadline - microtime(true);
        if ($remaining <= 0) {
            $erro = 'Tempo limite excedido (10s). Verifique se há loops infinitos.';
            proc_terminate($process, 9);
            break;
        }

        $changed = @stream_select($read, $write, $except, 0, 200000);

        if ($changed === false) break;

        foreach ($read as $stream) {
            $data = fread($stream, 8192);
            if ($stream === $pipes[1]) $output .= $data;
            if ($stream === $pipes[2]) $erro .= $data;
        }

        $status = proc_get_status($process);
        if (!$status['running']) break;
    }

    fclose($pipes[1]);
    fclose($pipes[2]);
    proc_close($process);
}

$elapsed = round(microtime(true) - $startTime, 3);

// Cleanup
@unlink($tmpFile);
// Também limpar possíveis arquivos gerados
$goFile = str_replace('.vrb', '_verbo.go', $tmpFile);
@unlink($goFile);
$binFile = str_replace('.vrb', '', $tmpFile);
@unlink($binFile);

// Response
echo json_encode([
    'output' => $output,
    'erro'   => $erro,
    'tempo'  => $elapsed,
], JSON_UNESCAPED_UNICODE);
