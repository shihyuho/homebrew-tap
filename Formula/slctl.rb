class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  version "2.2.6"
  
  if OS.mac?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-#{version}.tgz"
    sha256 "fc4185fbd34f36f994e9b5b55c6f4628d7210f9f16bb0397f8210c47f1238c6f"
  elsif OS.linux?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-linux-#{version}.tgz"
    sha256 "983ebd6b363b6cceec58cedd67b9d7fe9431d4473e705e4d9f7ac31728e2c5a9"
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
