<?php
/**
 * Verbo Portal — API de Transpilação
 *
 * Recebe código Verbo via POST JSON, transpila para Go,
 * e retorna o código Go gerado.
 *
 * Integração Yii2: Adaptar para um Controller/Action.
 *
 * Exemplo de request:
 *   POST /api/transpilar.php
 *   Content-Type: application/json
 *   {"codigo": "Exibir com (\"Olá!\")."}
 *
 * Exemplo de response:
 *   {"go_code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Olá!\")\n}\n", "erro": ""}
 */

header('Content-Type: application/json; charset=utf-8');
header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: POST, OPTIONS');
header('Access-Control-Allow-Headers: Content-Type');

if ($_SERVER['REQUEST_METHOD'] === 'OPTIONS') {
    http_response_code(204);
    exit;
}

if ($_SERVER['REQUEST_METHOD'] !== 'POST') {
    http_response_code(405);
    echo json_encode(['erro' => 'Método não permitido. Use POST.']);
    exit;
}

$input = json_decode(file_get_contents('php://input'), true);
if (!$input || empty($input['codigo'])) {
    http_response_code(400);
    echo json_encode(['erro' => 'Campo "codigo" é obrigatório.']);
    exit;
}

$codigo = $input['codigo'];

if (strlen($codigo) > 10000) {
    http_response_code(400);
    echo json_encode(['erro' => 'Código muito longo. Máximo 10.000 caracteres.']);
    exit;
}

// Caminho do binário
$verboBin = realpath(__DIR__ . '/../../build/verbo');
if (!$verboBin || !is_executable($verboBin)) {
    $verboBin = 'verbo';
}

// Criar arquivo temporário .vrb
$tmpDir = sys_get_temp_dir();
$tmpFile = tempnam($tmpDir, 'verbo_') . '.vrb';
file_put_contents($tmpFile, $codigo);

// Executar transpilação (compilar gera o .go)
$goFile = str_replace('.vrb', '_verbo.go', $tmpFile);

$projectRoot = realpath(__DIR__ . '/../../');
$cmd = 'cd ' . escapeshellarg($projectRoot) . ' && ' . escapeshellcmd($verboBin) . ' compilar ' . escapeshellarg($tmpFile) . ' 2>&1';
$output = shell_exec($cmd);

$go_code = '';
$erro = '';

if (file_exists($goFile)) {
    $go_code = file_get_contents($goFile);
} else {
    // Tentar pegar erro do output
    $erro = $output ?: 'Erro ao transpilar. Verifique a sintaxe do código.';
}

// Cleanup
@unlink($tmpFile);
@unlink($goFile);
$binFile = str_replace('.vrb', '', $tmpFile);
@unlink($binFile);

echo json_encode([
    'go_code' => $go_code,
    'erro' => $erro,
], JSON_UNESCAPED_UNICODE);
