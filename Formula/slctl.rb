class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  version "3.5.8"
  
  if OS.mac?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-#{version}.tgz"
    sha256 "e23e0a3f226e00c64dc755d2af12f12373667eada764de1481d4d5c61287136f"
  elsif OS.linux?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-linux-#{version}.tgz"
    sha256 "cf47879a34d115ee72f327dd12e17f47c7750a19fc22256de301e58e660b7bab"
  end

  depends_on :arch => :x86_64
  
  def install
    bin.install "slctl"
  end

  def caveats; <<~EOS
    To begin working with slctl, run the 'slctl init' command.
  EOS
  end
end
