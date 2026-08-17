cask "berth" do
  version "1.0.0"
  sha256 "3e5fa16021939cdbddbbee4c73bdc768fad2d79b21107fbbd29ec5b98d39b5ad"

  url "https://github.com/shihyuho/berth/releases/download/#{version}/Berth-arm64.zip"
  name "Berth"
  desc "Keep your macOS Dock on the display you choose"
  homepage "https://github.com/shihyuho/berth"

  depends_on macos: :ventura
  depends_on arch: :arm64

  app "Berth.app"

  caveats do
    unsigned_accessibility
  end
end
