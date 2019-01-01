class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  version "0.1.8"
  
  if OS.mac?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-#{version}.tgz"
  elsif OS.linux?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-linux-#{version}.tgz"
  end

  depends_on :arch => :x86_64
  depends_on "git"
  
  def install
    bin.install "slctl"
  end

  def caveats; <<~EOS
    To begin working with slctl, run the 'slctl init' command.
  EOS
  end
end
