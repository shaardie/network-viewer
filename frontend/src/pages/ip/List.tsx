import { useEffect, useState } from "react";

import { type IP } from "../../types/models";
import { Operations } from "../../components/Operations";

export function List() {
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

  return (
    <>
      <h1>IPs</h1>
      <div className="grid">
        <input
          type="search"
          name="search"
          placeholder="Search"
          aria-label="Search"
        />
      </div>
      <div className="container">
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
                    <Operations
                      id={ip.id}
                      type="ip"
                      onDelete={fetchIPs}
                    ></Operations>
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
