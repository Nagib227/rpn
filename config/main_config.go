package config

// Config содержит настройки ширины и высоты.
type Config struct {
    Width  int // Ширина
    Height int // Высота
}

// New создает новый экземпляр Config с значениями по умолчанию.
func New() *Config {
    return &Config{
        Width:  100, // Значение по умолчанию для Width
        Height: 50,  // Значение по умолчанию для Height
    }
}
