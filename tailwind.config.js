/** @type {import("tailwindcss").Config} */
module.exports = {
    content: [
        "!/vendor/**/*",
        "!/.devenv/**/*",
        "!/docs/**/*",
        "./**/*.{html,js,go,jsx,tsx}",
        // markdown dosyalari crash ettiriyor
        // ornek pkg/agents/kb/...
        "!./**/*.md",
    ]
};


