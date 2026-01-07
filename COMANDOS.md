# 📚 Documentación - Music Migration Tool

## 🚀 Inicio Rápido

### Iniciar la aplicación
```powershell
cd C:\Users\jorge.soriano\Documents\proyectos\music-migration
docker compose up -d
```

### Verificar que todo esté corriendo
```powershell
docker compose ps
```

### Acceder a la aplicación
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:9090
- **Health Check**: http://localhost:9090/health

---

## 🔧 Comandos de Docker

### Ver logs en tiempo real
```powershell
# Ver todos los logs
docker compose logs -f

# Ver logs solo del backend
docker compose logs -f backend

# Ver logs solo del frontend
docker compose logs -f frontend

# Ver logs de la base de datos
docker compose logs -f postgres
```

### Gestión de contenedores
```powershell
# Detener todos los contenedores
docker compose down

# Detener y eliminar volúmenes (borra la base de datos)
docker compose down -v

# Reiniciar todos los contenedores
docker compose restart

# Reiniciar solo un servicio
docker compose restart backend

# Ver estado de los contenedores
docker compose ps

# Ver recursos usados
docker stats
```

### Reconstruir después de cambios en el código
```powershell
# Reconstruir e iniciar
docker compose up -d --build

# Reconstruir solo el backend
docker compose up -d --build backend

# Reconstruir solo el frontend
docker compose up -d --build frontend
```

### Acceder a un contenedor
```powershell
# Acceder al contenedor del backend
docker exec -it music-migration-backend sh

# Acceder a la base de datos PostgreSQL
docker exec -it music-migration-db psql -U postgres -d music_migration

# Acceder a Redis
docker exec -it music-migration-redis redis-cli
```

---

## 🛠️ Desarrollo sin Docker

### Backend (Go)
```powershell
cd backend

# Instalar dependencias
go mod download

# Configurar variables de entorno
cp .env.example .env
# Edita .env con tus credenciales

# Ejecutar en desarrollo
go run cmd/api/main.go

# Compilar
go build -o music-migration.exe cmd/api/main.go

# Ejecutar tests
go test ./...

# Limpiar dependencias
go mod tidy
```

### Frontend (Next.js)
```powershell
cd frontend

# Instalar dependencias
npm install

# Configurar variables de entorno
cp .env.example .env.local
# Edita .env.local con NEXT_PUBLIC_API_URL=http://localhost:9090

# Ejecutar en desarrollo
npm run dev

# Compilar para producción
npm run build

# Ejecutar en producción
npm start

# Limpiar caché
Remove-Item -Recurse -Force .next
```

---

## 🐛 Solución de Problemas

### Error: "port is already in use"
```powershell
# Ver qué proceso usa el puerto
netstat -ano | findstr :9090

# Cambiar el puerto en docker-compose.yml
# Edita la sección backend > ports: "NUEVO_PUERTO:8080"
```

### Error: "no configuration file provided"
```powershell
# Asegúrate de estar en la carpeta correcta
cd C:\Users\jorge.soriano\Documents\proyectos\music-migration

# Verifica que existe docker-compose.yml
Get-ChildItem docker-compose.yml
```

### Los contenedores no inician
```powershell
# Ver logs detallados
docker compose up

# Limpiar todo y empezar de cero
docker compose down -v
docker system prune -a
docker compose up -d --build
```

### Base de datos no conecta
```powershell
# Verificar que PostgreSQL esté healthy
docker compose ps

# Ver logs de PostgreSQL
docker compose logs postgres

# Reiniciar PostgreSQL
docker compose restart postgres
```

### Frontend no carga
```powershell
# Verificar variables de entorno
docker compose logs frontend

# Reconstruir frontend
docker compose up -d --build frontend

# Acceder al contenedor y ver errores
docker exec -it music-migration-frontend sh
npm run build
```

---

## 📊 Base de Datos

### Conectar a PostgreSQL
```powershell
# Desde Docker
docker exec -it music-migration-db psql -U postgres -d music_migration

# Desde tu máquina (si tienes psql instalado)
psql -h localhost -p 5432 -U postgres -d music_migration
# Password: postgres
```

### Comandos SQL útiles
```sql
-- Ver todas las tablas
\dt

-- Ver estructura de una tabla
\d users

-- Ver todos los usuarios
SELECT * FROM users;

-- Ver migraciones
SELECT * FROM migrations ORDER BY created_at DESC;

-- Limpiar datos de prueba
TRUNCATE users, auth_tokens, migrations, track_mappings CASCADE;
```

---

## 🔄 Git y GitHub

### Comandos básicos
```powershell
# Ver estado
git status

# Agregar cambios
git add .

# Hacer commit
git commit -m "Descripción del cambio"

# Subir a GitHub
git push

# Ver historial
git log --oneline

# Ver diferencias
git diff
```

### Sincronizar con GitHub
```powershell
# Obtener últimos cambios
git pull

# Ver remotes
git remote -v

# Cambiar URL del repositorio
git remote set-url origin https://github.com/TU-USUARIO/music-migration.git
```

---

## 📦 Puertos Utilizados

| Servicio   | Puerto Host | Puerto Container | URL                      |
|------------|-------------|------------------|--------------------------|
| Frontend   | 3000        | 3000             | http://localhost:3000    |
| Backend    | 9090        | 8080             | http://localhost:9090    |
| PostgreSQL | 5432        | 5432             | localhost:5432           |
| Redis      | 6379        | 6379             | localhost:6379           |

---

## 🔐 Variables de Entorno

### Backend (.env)
```bash
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/music_migration?sslmode=disable
REDIS_URL=localhost:6379
SPOTIFY_CLIENT_ID=tu_client_id
SPOTIFY_CLIENT_SECRET=tu_client_secret
APPLE_TEAM_ID=tu_team_id
APPLE_KEY_ID=tu_key_id
JWT_SECRET=tu_secret_key
```

### Frontend (.env.local)
```bash
NEXT_PUBLIC_API_URL=http://localhost:9090
NEXT_PUBLIC_WS_URL=ws://localhost:9090/api/ws
```

---

## 📚 Recursos Adicionales

- **Spotify API Docs**: https://developer.spotify.com/documentation/web-api
- **Apple Music API**: https://developer.apple.com/documentation/applemusicapi
- **Go Documentation**: https://go.dev/doc/
- **Next.js Docs**: https://nextjs.org/docs
- **Fiber Framework**: https://docs.gofiber.io/
- **Docker Compose**: https://docs.docker.com/compose/

---

## 🎯 Próximos Pasos

1. **Obtener credenciales de Spotify**:
   - Ve a https://developer.spotify.com/dashboard
   - Crea una aplicación
   - Copia Client ID y Client Secret

2. **Obtener credenciales de Apple Music**:
   - Ve a https://developer.apple.com/
   - Crea un MusicKit identifier
   - Genera una private key

3. **Configurar OAuth completo**
4. **Implementar motor de matching**
5. **Agregar tests**
6. **Deploy a producción** (Railway, Fly.io, Vercel)

---

## 💡 Tips

- Siempre ejecuta `docker compose down` antes de cerrar tu computadora
- Usa `git commit` frecuentemente para no perder cambios
- Los logs son tu mejor amigo: `docker compose logs -f`
- Si algo no funciona, reconstruye: `docker compose up -d --build`
- Mantén tus credenciales en `.env` (nunca las subas a GitHub)

---

**¿Problemas?** Revisa los logs o abre un issue en GitHub.
