class Fx < Formula
  desc "Command-line JSON viewer 🔥"
  homepage "https://github.com/antonmedv/fx"

  depends_on "node" => :build

  def install
    system "npm", "install -g fx"
  end
end
