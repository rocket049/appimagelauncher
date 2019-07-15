# appimagelauncher 帮助 appimage 程序把自己加入启动菜单

调用： `func Create(desktopFile, iconFile string, force bool) error`

- `desktopFile`：`appimage` 的根目录中的 `.desktop` 文件。
- `iconFile`：`appimage` 的根目录中的图标文件。
- `force`：强制更新，如果 `force` 的值为 `false`，当启动器已经存在，并且比 `APPIMAGE` 更新时，不会重复创建起动器。

### 示例代码
```
import "gitee.com/rocket049/appimagelauncher"

appimagelauncher.Create("appimage-name.desktop", "icon-name.png", false)
```