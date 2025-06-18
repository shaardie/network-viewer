import { useEffect, useState } from "react";
import { IPBase } from "./IPBase";

import { type IP } from "../../types/models";

export function IPList() {
  const [ips, setIPs] = useState<IP[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchIPs = () => {
    setLoading(true);
    fetch("/api/v1/ip")
      .then((res) => res.json())
      .then((data) => {
        console.log("IP-Daten:", data);
        setIPs(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Fehler beim Laden der IPs:", err);
        setLoading(false);
      });
  };

  useEffect(() => {
    fetchIPs();
  }, []);

  const handleDelete = async (id: number) => {
    if (!confirm("IP wirklich löschen?")) return;

    const res = await fetch(`/api/v1/ip/${id}`, {
      method: "DELETE",
    });

    if (res.ok) {
      fetchIPs();
    } else {
      const errText = await res.text();
      alert("Löschen fehlgeschlagen: " + errText);
    }
  };

  return (
    <IPBase>
      {loading ? (
        <p>Loading...</p>
      ) : ips.length === 0 ? (
        <p>no IPs configured</p>
      ) : (
        <table>
          <thead>
            <tr>
              <th>IP</th>
              <th>Hostname</th>
              <th>Online</th>
              <th>RTT</th>
              <th>Comment</th>
              <th>Operations</th>
            </tr>
          </thead>
          <tbody>
            {ips.map((ip) => (
              <tr key={ip.id}>
                <td>{ip.ip}</td>
                <td>{ip.hostname}</td>
                <td>{ip.online ? "✅" : "❌"}</td>
                <td>{ip.rtt}</td>
                <td>{ip.comment}</td>
                <td>
                  <div className="grid">
                    <button
                      className="secondary"
                      onClick={() => handleDelete(ip.id)}
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
    </IPBase>
  );
}
