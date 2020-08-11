compress png and jpg

Example

We need open the file

```
file, err := os.OpenFile("12.png", os.O_RDONLY, 0444)
	if err != nil {
		fmt.Errorf("os.OpenFile: %s", err.Error())
	}
	defer file.Close()
```

```
PNGQuant(false, file, "14.png", 1, 0)
JPEGQuant(file, "20.jpeg", 1000, 0)
```