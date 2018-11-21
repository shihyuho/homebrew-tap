class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  url "https://github.com/softleader/slctl/releases/download/0.1.0/slctl-macos-.tgz"
  version "0.1.0"
  
  depends_on "git"

  def install
    bin.install "slctl"
  end

  def caveats; <<~EOS
    To begin working with slctl, run the 'slctl init' command.
  EOS
  end
end
