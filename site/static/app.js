async function callApi() {
  const out = document.getElementById('out');
  out.textContent = 'carregando...';
  try {
    const r = await fetch('/api/saude');
    const t = await r.text();
    out.textContent = t;
  } catch (e) {
    out.textContent = String(e);
  }
}

document.getElementById('btn').addEventListener('click', callApi);
