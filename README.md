## 介绍
+仅支持MacOS
MacOS全屏展示板，可以显示html。

## 使用
```bash
rm -r pkg/assets/web/*
cd web
yarn
yarn static
cp -r out/* ../pkg/assets/web/
cd ../
make build
./bin/*/wisper
```

## Thanks

Github: [https://github.com/progrium/macdriver](https://github.com/progrium/macdriver)

## License

MIT