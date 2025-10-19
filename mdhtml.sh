mkdir -p readme
curl -L -z "readme/sakura.css" "https://cdn.jsdelivr.net/npm/sakura.css/css/sakura.css"
for f in *.md; do
    echo "$f -> ${f%.md}.html"
    pandoc "$f" -s --embed-resources --standalone -c "readme/sakura.css" >"readme/${f%.md}.html"
done
