class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  version "3.6.0"
  
  if OS.mac?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-#{version}.tgz"
    sha256 "82a556880e0ae234222983d39f2c7364476cbfe22ffd0db648da2f883a3a841e"
  elsif OS.linux?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-linux-#{version}.tgz"
    sha256 "1a38dfef865288de33de7519b26ef89497e457289ed3bac27b83cdf96a985d26"
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
