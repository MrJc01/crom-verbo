# Acessibilidade — Boas Práticas

Este guia reúne boas práticas para tornar sites e aplicações acessíveis para o maior número de pessoas possível, incluindo pessoas com deficiências visuais, auditivas, motoras, cognitivas e pessoas usando tecnologias assistivas (leitores de tela, navegação por teclado, alto contraste, etc.).

## Princípios (POUR)

- **Perceptível**
  - Informação e componentes devem ser apresentados de forma que os usuários possam perceber.
- **Operável**
  - Interface e navegação devem ser utilizáveis por diferentes formas de interação (ex: teclado).
- **Compreensível**
  - Conteúdo e operação da interface devem ser previsíveis e fáceis de entender.
- **Robusto**
  - Conteúdo deve funcionar em navegadores e tecnologias assistivas diversas.

## HTML semântico

- Use elementos semânticos em vez de `div` genéricas.
  - `header`, `nav`, `main`, `footer`, `section`, `article`.
- Use títulos em hierarquia correta.
  - Um `h1` por página.
  - `h2`/`h3` para subseções, sem pular níveis.
- Para listas, use `ul`/`ol` e `li`.
- Para botões, use `button` (não `div` clicável).

## Imagens e mídia

- Toda imagem informativa deve ter `alt` descritivo.
  - Se for decorativa, use `alt=""`.
- Para vídeos:
  - Ofereça legendas.
  - Se possível, transcrição.
- Para áudio:
  - Transcrição.

## Formulários

- Use `label` associado ao `input`.
  - Preferir `for`/`id`.
- Use `fieldset` e `legend` para grupos de campos.
- Indique erros de forma clara:
  - Mensagens próximas ao campo.
  - Também sinalize via `aria-describedby`.
- Não dependa apenas de cor para indicar erro.

## Navegação por teclado

- Garanta que todos os controles interativos sejam acessíveis via teclado.
  - `Tab` deve alcançar todos os elementos.
  - Ordem de tabulação lógica.
- O foco deve ser visível.
  - Não remova `outline` sem substituto.
- Evite “armadilhas” de foco.

## Contraste e tipografia

- Garanta contraste suficiente entre texto e fundo.
  - Para texto normal, mirar em contraste equivalente a WCAG AA.
- Não use apenas cor para comunicar estado (erro, sucesso, seleção).
- Use tamanhos de fonte legíveis e espaçamento confortável.

## ARIA (use com moderação)

- Prefira HTML semântico; ARIA é para complementar.
- Use `aria-label`/`aria-labelledby` quando o texto visível não for suficiente.
- Use `role` apenas quando necessário.
- Mantenha estados sincronizados:
  - `aria-expanded`, `aria-selected`, `aria-pressed`.

## Mensagens e feedback

- Para notificações importantes, use regiões ao vivo quando apropriado:
  - `aria-live="polite"` para avisos não críticos.
  - `aria-live="assertive"` com cuidado.
- Mensagens devem ser objetivas e orientadas à ação.

## Links e botões

- Texto do link deve fazer sentido fora de contexto.
  - Evite “clique aqui”.
- Indique quando um link abre nova aba/janela.
- Botões devem descrever a ação.

## Componentes dinâmicos

- Para modais:
  - Trave foco dentro do modal.
  - Fechamento via `Esc`.
  - Retorne foco para o elemento que abriu.
- Para menus/dropdowns:
  - Suporte a teclado.
  - Estados ARIA adequados.

## Teste rápido (checklist)

- Navegue a página inteira usando apenas teclado.
- Teste com leitor de tela (quando possível).
- Verifique contraste.
- Aumente zoom para 200%.
- Desative CSS e veja se o HTML ainda faz sentido.

## Referências

- WCAG (Web Content Accessibility Guidelines)
- WAI-ARIA Authoring Practices
