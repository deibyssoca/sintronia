# 🌱 Diagrama Mermaid - Modelo de Datos Sintropia

## 📊 Diagrama de Entidad-Relación (Mermaid)

```mermaid
erDiagram
    LOCATION {
        uint id PK
        string name UK "UNIQUE"
        string notes
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at "soft delete"
    }
    
    ARRANGEMENT {
        uint id PK
        uint location_id FK
        string name
        string type "linea, isla, gremio"
        float64 length "solo para líneas"
        float64 width "solo para líneas"
        float64 diameter "solo para islas"
        string soil_type
        string planting_mode
        string notes
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at "soft delete"
    }
    
    PLANT {
        uint id PK
        string name
        string scientific
        string stratum "emergente, alto, medio, bajo, rastrero, trepador, raiz"
        string function "fijador_nitrogeno, acumulador_dinamico, etc"
        string succession_stage "placenta, pionera, secundaria, climax"
        string external_id UK "UNIQUE - referencia API externa"
        bool desired "lista de plantas deseadas"
        string notes
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at "soft delete"
    }
    
    PLANTING {
        uint id PK
        uint arrangement_id FK
        uint plant_id FK
        int quantity "cantidad de plantas"
        string status "planeada, germinacion, plantada, etc"
        string position "descripción de posición"
        string notes
        timestamp planted_at "fecha de plantación"
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at "soft delete"
    }
    
    %% Relaciones
    LOCATION ||--o{ ARRANGEMENT : "tiene múltiples"
    ARRANGEMENT ||--o{ PLANTING : "contiene"
    PLANT ||--o{ PLANTING : "se planta en"
```

## 🔄 Diagrama de Flujo de Proceso

```mermaid
graph TD
    A[Crear Location] --> B[Crear Arrangement]
    C[Agregar Plant al catálogo] --> D[Crear Planting]
    B --> D
    
    D --> E{Qué tipo de Arrangement?}
    E -->|Línea| F[Calcular área: length × width]
    E -->|Isla| G[Calcular área: π × diameter/2²]
    E -->|Gremio| H[Área variable]
    
    F --> I[Calcular densidad]
    G --> I
    H --> I
    
    I --> J[Generar reportes]
    
    style A fill:#e1f5fe
    style B fill:#f3e5f5
    style C fill:#e8f5e8
    style D fill:#fff3e0
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

## 🚀 Para visualizar estos diagramas:

### **Opción 1: GitHub/GitLab**
- Los archivos `.md` con Mermaid se renderizan automáticamente

### **Opción 2: Mermaid Live Editor**
- Ir a: https://mermaid.live/
- Copiar y pegar el código Mermaid

### **Opción 3: VS Code**
- Instalar extensión "Mermaid Preview"
- Abrir archivo `.md` y usar preview

### **Opción 4: Generar PNG con Python**
```bash
pip install graphviz
python docs/generate-diagram.py
```