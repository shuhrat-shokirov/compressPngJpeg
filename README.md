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

If your OC Windows: need install MINGW

 ```html
 i remove my mingw 32 bit installtion and using mingw 64 installation from souceforge site here

https://sourceforge.net/projects/mingw-w64/

i follow default installation setup and it install mingw in this folder

C:\Program Files\mingw-w64\x86_64-6.2.0-posix-seh-rt_v5-rev0\mingw64\

so put C:\Program Files\mingw-w64\x86_64-6.2.0-posix-seh-rt_v5-rev0\mingw64\bin folder in PATH directory and test with gcc command from cmd.
```
