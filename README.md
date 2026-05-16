# AIDD-repository
Un MCP para correr en local, que incluye las bases para hacer **Spec Driven Development** seguro de forma y agnostica ( **sin vendor lock-in** ), permitiendo cambiar de coding agent en cualquier momento sin perder tu progreso o sufrir tenidiendo que hacer un setup completo al cambiar ( **Bueno, solo agregar el config del MCP UwU** ).

### Como instalar

Copia, pega y corre el siguiente comando en tu terminal:
```
curl -fsSL https://raw.githubusercontent.com/gabriel96x1/AIDD-repository/main/install.sh | bash
```

Para las coding conventions usar en el repo del proyecto la herramienta de Autoskill:

https://github.com/midudev/autoskills/tree/main

### Como hacer setup del MCP VS Code:

En el root de tu proyecto ve a .vscode/mcp.json y pega lo siguiente:
```
{
	"servers": {
		"aidd-server": {
			"type": "stdio",
			"command": "./aidd-mcp-server",
			"args": []
		}
	},
	"inputs": []
}
```

### Como usar el MCP para SDD:

[!WARNING]
> Para hacer setup de tu proyecto para trabajar con esta herramiente de SDD debes comenzar con el siguiente PROMPT: 
```
Usando el mcp de aidd-server haz setup de este proyecto para trabajar con Spec Driven Development"
```

#### Iniciar feature
```
"Quiero [feature]. Empieza el proceso SDD."
```
#### Avanzar fase
```
"Aprobado. Procede a [design/tasks/execute]."
```

#### Corrección en gate
```
"[AC-N] no es correcto. Debe ser: [descripción]. Actualiza el spec."
```

#### Cambio pequeño sin spec completo
```
"Esto es un chore/bugfix, no necesita spec. Hazlo y documenta en evidence.md."
```

#### Consultar estado
```
"¿En qué gate estamos para [feature-slug]? Muéstrame el spec actual."
```

#### Replanning
```
"La feature X cambió de scope. Abre una sesión de replanning."
```

#### Contexto persistente
```
"Decidimos [X]. Guárdalo en [tech-stack/mission/roadmap].md"
```