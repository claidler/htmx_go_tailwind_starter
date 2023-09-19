air -c ./.air.toml & \
npx tailwind \
  -i 'styles.css' \
  -o 'public/styles.css' \
  --watch & \
browser-sync start \
  --files 'public/**/*.html, public/**/*.css' \
  --port 3001 \
  --proxy 'localhost:3000' \
  --middleware 'function(req, res, next) { \
    res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
    return next(); \
  }'