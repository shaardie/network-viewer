import { type ReactNode } from "react";

export function Layout({ children }: { children: ReactNode }) {
  return (
    <>
      <header className="container">
        <nav>
          <ul>
            <li>
              <strong>
                <a
                  href="/"
                  style={{
                    all: "unset",
                    cursor: "pointer",
                  }}
                >
                  Network Viewer
                </a>
              </strong>
            </li>
          </ul>
          <ul>
            <li>
              <a href="/subnet">Subnets</a>
            </li>
            <li>
              <a href="/ip">IPs</a>
            </li>
          </ul>
        </nav>
      </header>

      <main className="container">{children}</main>

      <footer className="container">
        <hr />
        <small>
          ©2025 Network Viewer –{" "}
          <a
            href="https://github.com/shaardie/network-viewer"
            target="_blank"
            rel="noopener noreferrer"
          >
            GitHub
          </a>
        </small>
      </footer>
    </>
  );
}
