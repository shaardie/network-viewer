import { useEffect, useState, useRef } from "react";
import { useParams } from "react-router-dom";
import { formatDurationNs, formatDateTime } from "../../lib/lib.ts";
import { type Subnet } from "../../types/models.ts";
import { Operations } from "../../components/subnet/Operations.tsx";

export function Show() {
  const { id } = useParams<{ id: string }>();
  const [subnet, setSubnet] = useState<Subnet | null>(null);
  const [error, setError] = useState("");
  const dialogRef = useRef<HTMLDialogElement>(null);

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

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = new FormData(e.currentTarget);

    const hours = Number(form.get("scanner_hours") || 0);
    const minutes = Number(form.get("scanner_minutes") || 0);
    const seconds = Number(form.get("scanner_seconds") || 0);

    const totalSeconds = hours * 3600 + minutes * 60 + seconds;
    const durationNs = totalSeconds * 1_000_000_000; // Go erwartet int64 ns

    const payload = {
      subnet: form.get("subnet"),
      comment: form.get("comment"),
      scanner_enabled: form.get("scanner_enabled") === "true",
      scanner_interval: durationNs,
    };

    const res = await fetch("/api/v1/subnet", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });

    if (!res.ok) {
      const errText = await res.text();
      setError(errText);
    } else {
      window.location.href = "/subnet";
    }
  };

  return (
    <>
      <h1>Subnet: {subnet.subnet}</h1>
      <div className="grid">
        <Operations subnet={subnet}></Operations>
      </div>
      <dialog ref={dialogRef}>
        <article>
          <div className="container">
            {error && <strong style={{ color: "red" }}>{error}</strong>}
            <form onSubmit={handleSubmit}>
              <fieldset>
                <label>
                  Subnet
                  <input
                    type="text"
                    name="subnet"
                    defaultValue={subnet.subnet}
                    placeholder="192.168.0.0/24"
                    required
                  />
                </label>
                <label>
                  Comment
                  <input
                    type="text"
                    name="comment"
                    defaultValue={subnet.comment}
                    placeholder="Comment"
                  />
                </label>
              </fieldset>
              <fieldset>
                <legend>Scanner</legend>
                <label>
                  <input
                    type="checkbox"
                    name="scanner_enabled"
                    defaultValue="true"
                    defaultChecked
                  />
                  Activate
                </label>
                <div className="grid">
                  <label>
                    Hours
                    <input
                      type="number"
                      name="scanner_hours"
                      defaultValue={0}
                      min={0}
                    />
                  </label>
                  <label>
                    Minutes
                    <input
                      type="number"
                      name="scanner_minutes"
                      defaultValue={5}
                      min={0}
                      max={59}
                    />
                  </label>
                  <label>
                    Seconds
                    <input
                      type="number"
                      name="scanner_seconds"
                      defaultValue={0}
                      min={0}
                      max={59}
                    />
                  </label>
                </div>
              </fieldset>
              <button type="submit">Edit Subnet</button>
            </form>
          </div>
        </article>
      </dialog>
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
