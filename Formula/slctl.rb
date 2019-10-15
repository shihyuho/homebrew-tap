class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  version "3.8.0"
  
  if OS.mac?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-#{version}.tgz"
    sha256 "643433856f1613aaec8f8d3a2fbf6aadfb0a10c423ce697d01c4c2bd2bf3dabe"
  elsif OS.linux?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-linux-#{version}.tgz"
    sha256 "7290d23180859f89a6a57f20b0312a637c5d4978036d9857a8f2941eb420fad3"
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
