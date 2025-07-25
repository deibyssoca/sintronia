# 🌱 Diagrama Mermaid - Modelo de Datos Sintropia

## 📊 Diagrama de Entidad-Relación (Mermaid)

```mermaid
erDiagram
    %% Entidades principales
    SITE {
        uint    id PK
        string  name
        float   area_m2     "área total calculada"
        float   length_m
        float   width_m
        string  notes
        timestamp created_at
        timestamp updated_at
    }

    PLANTATION {
        uint    id PK
        uint    site_id FK
        string  name
        float   area_m2     "área definida o calculada"
        string  notes
        timestamp created_at
        timestamp updated_at
    }

    PLOT {
        uint    id PK
        uint    plantation_id FK
        enum    plot_type   "line, island, guild"
        float   length_m    "para líneas"
        float   width_m     "para líneas"
        float   diameter_m  "para islas"
        string  geometry    "GeoJSON opcional (PostGIS)"
        timestamp created_at
        timestamp updated_at
    }

    PLANT_SPECIES {
        uint    id PK
        string  common_name
        string  scientific_name
        enum    stratum            "emergente, alto, medio, ..."
        enum    function_ecol      "fijador, acumulador, ..."
        enum    succession_stage   "pionera, secundaria, ..."
        string  external_ref       "ID en API externa"
        timestamp created_at
        timestamp updated_at
    }

    PLANT_INSTANCE {
        uint    id PK
        uint    plot_id FK
        uint    species_id FK
        int     quantity
        enum    role               "objetivo, servicio, acompañante"
        enum    status             "planeada, germinada, plantada, desarrollada, muerta, eliminada, podada"
        string  position           "GeoJSON o descripción"
        date    planted_at
        timestamp created_at
        timestamp updated_at
    }

    SUGGESTION_TEMPLATE {
        uint    id PK
        uint    plantation_id FK
        string  name
        string  description
        json    rules              "reglas de densidad, estrato, sucesión"
        timestamp created_at
        timestamp updated_at
    }

    %% Relaciones
    SITE          ||--o{ PLANTATION          : "tiene"
    PLANTATION   ||--o{ PLOT                 : "define"
    PLOT         ||--o{ PLANT_INSTANCE       : "contiene"
    PLANT_SPECIES||--o{ PLANT_INSTANCE       : "especie de"
    PLANTATION   ||--o{ SUGGESTION_TEMPLATE  : "ofrece"
```

## 🔄 Diagrama de Flujo de Proceso

```mermaid
flowchart TD
    A[Crear Sitio ] --> B[Crear Plantación ]
    B --> C[Agregar Parcelas sintrópicas]
    C --> D[Insertar Plantas a mano ]
    C --> E[Aplicar Sugerencia automática]
    E --> F[Generar distribución sugerida]
    F --> G[Insertar plantas sugeridas]
    D --> H[Visualizar o Editar plantación]
    G --> H
    H --> I[Guardar o Ejecutar siembra]

    subgraph Base de Datos
        A
        B
        C
        D
        E
        G
    end

    style A fill:#bbf,stroke:#333,stroke-width:1px
    style B fill:#bbf,stroke:#333,stroke-width:1px
    style C fill:#cfc,stroke:#333,stroke-width:1px
    style D fill:#ffd,stroke:#333,stroke-width:1px
    style E fill:#ffd,stroke:#333,stroke-width:1px
    style F fill:#eef,stroke:#333,stroke-width:1px
    style G fill:#bbf,stroke:#333,stroke-width:1px
    style H fill:#ccc,stroke:#333,stroke-width:1px
    style I fill:#faa,stroke:#333,stroke-width:1px
```

## 🎯 Diagrama de Estados de Plantación

```mermaid
stateDiagram-v2
    [*] --> Planeada
    Planeada --> Germinacion : iniciar proceso
    Germinacion --> Plantula : germina exitosamente
    Germinacion --> Muerta : falla germinación
    Plantula --> Plantada : trasplantar al campo
    Plantula --> Muerta : falla en plantula
    Plantada --> Establecida : se adapta al terreno
    Plantada --> Muerta : no se adapta
    Establecida --> Productiva : comienza producción
    Establecida --> Dormante : entra en dormancia
    Productiva --> Dormante : ciclo estacional
    Dormante --> Productiva : sale de dormancia
    Productiva --> Muerta : fin de ciclo
    Dormante --> Muerta : no sale de dormancia
    Muerta --> [*]
```

## 📋 Constantes del Sistema

```mermaid
mindmap
  root((Sintropia))
    Estratos de Vegetación
      Emergente mayor 25m
      Alto 15-25m
      Medio 5-15m
      Bajo 1-5m
      Rastrero menor 1m
      Trepador
      Raíz
    Etapas Sucesionales
      Placenta
      Pionera
      Secundaria
      Clímax
    Funciones Ecológicas
      Fijador Nitrógeno
      Acumulador Dinámico
      Cobertura Suelo
      Cortaviento
      Polinizador
      Control Plagas
      Aireación Suelo
      Regulación Agua
      Producción Biomasa
      Alimentario
      Medicinal
      Maderable
      Fibra
      Ornamental
    Tipos de Lecho
      Línea
      Isla
      Gremio
```
