import { type ReactNode } from "react";

export function IPBase({ children }: { children: ReactNode }) {
  return (
    <>
      <h1>IPs</h1>
      <nav>
        <ul>
          <li>
            <input
              type="search"
              name="search"
              placeholder="Search"
              aria-label="Search"
            />
          </li>
        </ul>
      </nav>
      <section className="container">{children}</section>
    </>
  );
}
