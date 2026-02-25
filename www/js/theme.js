/**
 * Verbo Portal — Dark/Light Theme Toggle
 * Persiste em localStorage e respeita prefers-color-scheme
 */
(function () {
    const STORAGE_KEY = 'verbo-theme';

    function getPreferred() {
        const stored = localStorage.getItem(STORAGE_KEY);
        if (stored) return stored;
        return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    }

    function apply(theme) {
        document.documentElement.classList.toggle('dark', theme === 'dark');
        localStorage.setItem(STORAGE_KEY, theme);
        // Update toggle icons
        document.querySelectorAll('[data-theme-icon]').forEach(el => {
            el.style.display = el.dataset.themeIcon === theme ? 'none' : 'block';
        });
    }

    function toggle() {
        const current = document.documentElement.classList.contains('dark') ? 'dark' : 'light';
        apply(current === 'dark' ? 'light' : 'dark');
    }

    // Apply on load
    apply(getPreferred());

    // Listen for system changes
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', e => {
        if (!localStorage.getItem(STORAGE_KEY)) {
            apply(e.matches ? 'dark' : 'light');
        }
    });

    // Export toggle
    window.verboTheme = { toggle, apply, getPreferred };
})();
