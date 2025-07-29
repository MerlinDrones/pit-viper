# pit-viper

A reusable, modular application configuration system for Go the leverages and uses [Viper](https://github.
com/spf13/viper), supporting TOML, environment overrides, and per-module defaults.

Additionally, `pit-viper` is designed for hexagonal or modular applications where configuration is initialized once at the application boundary, but each module is responsible for defining and registering its own defaults.

---

## 🧩 Features

- Per-module config registration via interfaces
- Automatic TOML generation from defaults
- Optional config file loading (with fallback to defaults)
- Environment variable overrides (`MYAPP_DB_DSN`, etc.)
- Code generation for registering module configs
- Reusable and composable within any Go project

---

## 🏗️ Project Structure

```plaintext
.
├── pkg/
│   ├── config/            # Core config loading logic
│   └── app.go             # APP_NAME / ENV_PREFIX constants
├── tools/
│   └── gen_config_list/   # Generates modules.gen.go with DefaultModules
├── examples/              # Optional examples using pit-viper
└── README.md
````

---

## 🚀 Usage

### 1. Install

```bash
  go get github.com/your-org/pit-viper@latest
```

---

### 2. Initialize Configuration

In your app’s `main()` or `init()`:

```go
import "your-org/pit-viper/pkg/config"

func main() {
	if err := config.Init(""); err != nil {
		log.Fatal(err)
	}
}
```

This will:

* Register all module defaults
* Attempt to load a `config.toml`
* Fallback to defaults + environment overrides

---

### 3. Load Configuration Values

```go
port := viper.GetInt("server.port")
dsn := viper.GetString("db.dsn")
```

Or use strongly-typed modules:

```go
cfg := db.NewDBConfig()
fmt.Println(cfg.DSN)
```

---

## 🛠 Customizing `pit-viper` for Reuse in Other Projects

### 🔁 1. Copy the Code

Copy `pkg/config` and `pkg/app.go` into your own project (e.g. `myapp/pkg/config`).

---

### ✏️ 2. Update `pkg/app.go`

```go
const (
	APP_NAME       = "myapp"
	APP_ENV_PREFIX = "MYAPP"
)
```

---

### 🧩 3. Create Your Own Modules

Each module should implement:

```go
type MyConfig struct { ... }

func (c *MyConfig) RegisterDefaults(v *viper.Viper) { ... }
func (c *MyConfig) String() string { ... }
```

---

### ⚙️ 4. Wire Your Modules

Use `tools/gen_config_list` or manually update `DefaultModules`:

```go
config.DefaultModules = []config.IConfig{
	db.NewDBConfig(),
	server.NewServerConfig(),
}
```

---

## ⚙️ Application Constants (`pkg/app.go`)

```go
const (
	APP_NAME       = "pit-viper"
	APP_ENV_PREFIX = "PIT"
)
```

These values drive:

* The default config file name: `pit-viper.toml`
* Environment var overrides: `PIT_DB_DSN`, etc.

---

## 🧪 Testing

```bash
  go test ./... -v
```

Includes:

* Tests for default registration
* Config loading from TOML and env vars
* Config file fallback behavior

---

## 📝 Example Config (`config.toml`)

```toml
[db]
dsn = "postgres://localhost/myapp"

[server]
port = 8080
```

---

## 📄 License

MIT