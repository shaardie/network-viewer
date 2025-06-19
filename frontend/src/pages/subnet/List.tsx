import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import { type Subnet } from "../../types/models.ts";
import { formatDurationNs, formatDateTime } from "../../lib/lib.ts";

import { Operations } from "../../components/subnet/Operations.tsx";

export function List() {
  const [subnets, setSubnets] = useState<Subnet[]>([]);
  const [loading, setLoading] = useState(true);

  const navigate = useNavigate();

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

  return (
    <>
      <h1>Subnets</h1>
      <div className="grid">
        <button onClick={() => navigate(`/subnet/create`)}>
          Create Subnet
        </button>
      </div>
      <div className="container">
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
                <tr
                  key={s.id}
                  style={{ cursor: "pointer" }}
                  onClick={() => navigate(`/subnet/${s.id}`)}
                >
                  <td>{s.subnet}</td>
                  <td>{s.scanner_enabled ? "✅" : "❌"}</td>
                  <td>{formatDurationNs(s.scanner_interval)}</td>
                  <td>{formatDateTime(s.last_scan)}</td>
                  <td>{s.comment}</td>
                  <td>
                    <Operations subnet={s}></Operations>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </>
  );
}
