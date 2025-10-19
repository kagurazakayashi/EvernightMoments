MD readme
curl -L -z "readme/sakura.css" "https://cdn.jsdelivr.net/npm/sakura.css/css/sakura.css"
for %%f in (*.md) do (
    pandoc "%%f" -s --embed-resources --standalone -c "readme\sakura.css" >"readme\%%~nf.html"
)
