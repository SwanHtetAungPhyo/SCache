class Scache < Formula
  desc "A simple caching system inspired by Redis, written in Go"
  homepage "https://github.com/SwanHtetAungPhyo/SCache"
  url "https://github.com/SwanHtetAungPhyo/Scache/releases/download/v1.0.0/scache-v1.0.0.tar.gz"
  sha256 "3546a2e2211fce317e7b89cc06e6e696c6b0ca21d50f625e38b76008f8022bc5"  # Replace with the actual SHA256
  license "MIT"

  def install
    bin.install "scache"  # Ensure "scache" is the name of the executable in the tarball
  end

  test do
    system "#{bin}/scache", "--version"
  end
end
