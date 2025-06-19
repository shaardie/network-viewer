import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { formatDurationNs, formatDateTime } from "../../lib/lib.ts";
import { type Subnet } from "../../types/models.ts";
import { Operations } from "../../components/Operations.tsx";

export function Show() {
  const { id } = useParams<{ id: string }>();
  const [subnet, setSubnet] = useState<Subnet | null>(null);
  const [error, setError] = useState("");

  useEffect(() => {
    fetch(`/api/v1/subnet/${id}`)
      .then((res) => {
        if (!res.ok) throw new Error("Error while loading");
        return res.json();
      })
      .then(setSubnet)
      .catch((err) => setError(err.message));
  }, [id]);

  if (error) return <p style={{ color: "red" }}>{error}</p>;
  if (!subnet) return <p>loading...</p>;

  return (
    <>
      <h1>Subnet: {subnet.subnet}</h1>
      <div className="grid">
        <Operations id={subnet.id} type="subnet"></Operations>
      </div>
      <table>
        <tbody>
          <tr>
            <th>ID</th>
            <td>{subnet.id}</td>
          </tr>
          <tr>
            <th>Subnet</th>
            <td>{subnet.subnet}</td>
          </tr>
          <tr>
            <th>Scanner enabled</th>
            <td>{subnet.scanner_enabled ? "✅" : "❌"}</td>
          </tr>
          <tr>
            <th>Scanner Interval</th>
            <td>{formatDurationNs(subnet.scanner_interval)}</td>
          </tr>
          <tr>
            <th>Last Scan</th>
            <td>{formatDateTime(subnet.last_scan)}</td>
          </tr>
          <tr>
            <th>Comment</th>
            <td>{subnet.comment}</td>
          </tr>
        </tbody>
      </table>
    </>
  );
}
