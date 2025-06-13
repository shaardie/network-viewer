import { useEffect, useState } from "react";
import { SubnetBase } from "./SubnetBase";

type Subnet = {
  id: number;
  subnet: string;
  scanner_enabled: boolean;
  scanner_interval: number;
  last_scan: string;
  comment: string;
};

export function SubnetList() {
  const [subnets, setSubnets] = useState<Subnet[]>([]);
  const [loading, setLoading] = useState(true);

  function formatDurationNs(nanoseconds: number): string {
    const totalSeconds = Math.floor(nanoseconds / 1_000_000_000);
    const hours = Math.floor(totalSeconds / 3600);
    const minutes = Math.floor((totalSeconds % 3600) / 60);
    const seconds = totalSeconds % 60;

    const parts = [];
    if (hours) parts.push(`${hours}h`);
    if (minutes || hours) parts.push(`${minutes}m`);
    parts.push(`${seconds}s`);
    return parts.join(" ");
  }
  const fetchSubnets = () => {
    setLoading(true);
    fetch("/api/v1/subnet")
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setSubnets(data);
        setLoading(false);
      });
  };

  useEffect(() => {
    fetchSubnets();
  }, []);

  const handleDelete = async (id: number) => {
    if (!confirm("Really delete this subnet?")) return;

    const res = await fetch(`/api/v1/subnet/${id}`, {
      method: "DELETE",
    });

    if (res.ok) {
      fetchSubnets();
    } else {
      const errText = await res.text();
      alert("Delete failed: " + errText);
    }
  };

  return (
    <SubnetBase>
      {loading ? (
        <p>Loading...</p>
      ) : subnets.length === 0 ? (
        <p>No Subnets configured</p>
      ) : (
        <table>
          <thead>
            <tr>
              <th>Subnet</th>
              <th>Scanner enabled</th>
              <th>Scanner Interval</th>
              <th>Last Scan</th>
              <th>Comment</th>
              <th>Operations</th>
            </tr>
          </thead>
          <tbody>
            {subnets.map((s) => (
              <tr key={s.id}>
                <td>{s.subnet}</td>
                <td>{s.scanner_enabled ? "✅" : "❌"}</td>
                <td>{formatDurationNs(s.scanner_interval)}</td>
                <td>{s.last_scan}</td>
                <td>{s.comment}</td>
                <td>
                  <div className="grid">
                    <button
                      className="secondary"
                      onClick={() => handleDelete(s.id)}
                    >
                      Delete
                    </button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </SubnetBase>
  );
}
