class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  url "https://github.com/softleader/homebrew-tap/releases/download/slctl/slctl-darwin-0.1.8.tgz"
  version "0.1.8"
  
  depends_on "git"

  def install
    bin.install "slctl"
  end

  def caveats; <<~EOS
    To begin working with slctl, run the 'slctl init' command.
  EOS
  end
end
