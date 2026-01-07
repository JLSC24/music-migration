# 🎵 Music Migration Tool

Migra tus playlists y biblioteca musical entre **Spotify** y **Apple Music** de forma rápida y eficiente.

## 🚀 Características

-   ✅ Migración bidireccional: Spotify → Apple Music y Apple Music → Spotify
-   ✅ Autenticación OAuth2 segura
-   ✅ Motor de matching inteligente con múltiples criterios (ISRC, metadata)
-   ✅ Procesamiento paralelo de miles de canciones
-   ✅ Progreso en tiempo real con WebSockets
-   ✅ Reporte detallado de éxitos y fallos
-   ✅ Interfaz moderna y responsive

## 🛠️ Stack Tecnológico

### Backend

-   **Go 1.21+** - Performance y concurrencia
-   **Fiber** - Framework web rápido
-   **PostgreSQL** - Base de datos
-   **Redis** - Cache y rate limiting

### Frontend

-   **Next.js 14** - React framework
-   **TypeScript** - Type safety
-   **Tailwind CSS** - Estilos
-   **TanStack Query** - Data fetching

## 📋 Prerequisitos

-   Go 1.21 o superior
-   Node.js 18+ y npm/yarn/pnpm
-   PostgreSQL 14+
-   Redis 7+
-   Docker & Docker Compose (opcional pero recomendado)

## 🏃 Quick Start

### Con Docker (Recomendado)

```bash
# 1. Clonar el repositorio
git clone https://github.com/tu-usuario/music-migration.git
cd music-migration

# 2. Configurar variables de entorno
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env.local
# Edita los archivos .env con tus credenciales

# 3. Iniciar todos los servicios
docker-compose up -d

# 4. Acceder a la aplicación
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080
```

### Sin Docker

#### Backend

```bash
cd backend

# Instalar dependencias
go mod download

# Configurar variables de entorno
cp .env.example .env
# Edita .env con tus credenciales

# Ejecutar migraciones
go run cmd/migrate/main.go

# Iniciar servidor
go run cmd/api/main.go
```

#### Frontend

```bash
cd frontend

# Instalar dependencias
npm install

# Configurar variables de entorno
cp .env.example .env.local
# Edita .env.local con la URL del backend

# Iniciar en desarrollo
npm run dev
```

## 🔑 Configuración de APIs

### Spotify API

1. Ve a [Spotify Developer Dashboard](https://developer.spotify.com/dashboard)
2. Crea una nueva aplicación
3. Obtén tu `Client ID` y `Client Secret`
4. Añade `http://localhost:8080/api/auth/spotify/callback` a Redirect URIs

### Apple Music API

1. Ve a [Apple Developer Portal](https://developer.apple.com/)
2. Crea un MusicKit identifier
3. Genera una Private Key (.p8 file)
4. Obtén tu `Team ID`, `Key ID` y guarda el archivo `.p8`

## 📁 Estructura del Proyecto

```
music-migration/
├── backend/                 # API Go
│   ├── cmd/
│   │   └── api/            # Entry point
│   ├── internal/
│   │   ├── handlers/       # HTTP handlers
│   │   ├── services/       # Business logic
│   │   ├── models/         # Data models
│   │   ├── middleware/     # Middleware
│   │   └── database/       # DB connection
│   ├── pkg/
│   │   ├── spotify/        # Spotify client
│   │   ├── apple/          # Apple Music client
│   │   └── matcher/        # Matching engine
│   └── go.mod
│
├── frontend/               # Next.js app
│   ├── src/
│   │   ├── app/           # Pages (App Router)
│   │   ├── components/    # React components
│   │   ├── lib/           # Utilities
│   │   └── types/         # TypeScript types
│   └── package.json
│
└── docker-compose.yml
```

## 🔄 Flujo de Migración

1. **Autenticación**: Usuario autoriza acceso a ambas plataformas
2. **Selección**: Elige playlists o biblioteca completa
3. **Extracción**: Sistema obtiene todas las canciones
4. **Matching**: Motor busca coincidencias en la plataforma destino
5. **Creación**: Genera playlists/añade a biblioteca
6. **Reporte**: Muestra resultados y canciones no encontradas

## 🧪 Testing

```bash
# Backend tests
cd backend
go test ./... -v

# Frontend tests
cd frontend
npm run test
```

## 📝 API Documentation

Una vez iniciado el backend, accede a:

-   Swagger UI: http://localhost:8080/swagger
-   API Docs: http://localhost:8080/api/docs

## 🤝 Contribuir

Las contribuciones son bienvenidas! Por favor:

1. Fork el proyecto
2. Crea una rama (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

MIT License - ve [LICENSE](LICENSE) para más detalles

## 👤 Autor

Tu Nombre - [@JLSC24](https://github.com/JLSC24)
