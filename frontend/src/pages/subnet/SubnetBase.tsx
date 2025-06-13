import { type ReactNode } from "react";
import { Link } from "react-router-dom";

export function SubnetBase({ children }: { children: ReactNode }) {
  return (
    <>
      <h1>Subnets</h1>
      <nav>
        <ul>
          <li>
            <Link to="/subnet">List Subnets</Link>
          </li>
          <li>
            <Link to="/subnet/create">Create new Subnet</Link>
          </li>
        </ul>
      </nav>
      <section className="container">{children}</section>
    </>
  );
}
