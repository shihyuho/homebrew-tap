# Homebrew Tap

Shihyu 個人維護的 Homebrew tap。各套件的功能、需求與專屬安裝說明，請以其專案 README 為準。

## 使用

先加入 tap：

```sh
brew tap shihyuho/tap
```

將範例中的名稱替換成實際套件名稱，再安裝 Formula 或 Cask：

```sh
brew install shihyuho/tap/FORMULA_NAME
brew install --cask shihyuho/tap/CASK_NAME
```

更新 Homebrew 資訊，再更新指定套件：

```sh
brew update
brew upgrade FORMULA_NAME
brew upgrade --cask CASK_NAME
```
