### Hexlet tests and linter status:

[![Actions Status](https://github.com/Subwme/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Subwme/go-project-242/actions)
[![Go](https://github.com/Subwme/go-project-242/actions/workflows/go.yml/badge.svg)](https://github.com/Subwme/go-project-242/actions/workflows/go.yml)

asciinema - https://asciinema.org/a/MGnF404sXkpUBZ7o

# hexlet-path-size

Утилита для подсчета размера файлов и директорий.

## Установка

```bash
make install
```

### Использование

```bash
# Размер файла
./bin/hexlet-path-size path/to/file.txt

# Размер директории
./bin/hexlet-path-size path/to/dir

# Человекочитаемый формат
./bin/hexlet-path-size path/to/dir --human

# Включить скрытые файлы
./bin/hexlet-path-size path/to/dir --all

# Рекурсивно обойти поддиректории
./bin/hexlet-path-size path/to/dir --recursive

# Комбинация флагов
./bin/hexlet-path-size path/to/dir -H -a -r
```

### Флаги
```bash
--human, -H - человекочитаемый формат (KB, MB, GB, и т.д.)
--all, -a - включить скрытые файлы и директории
--recursive, -r - рекурсивно обойти все поддиректории
```