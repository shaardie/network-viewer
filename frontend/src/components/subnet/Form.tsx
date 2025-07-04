import { useState } from "react";
import { type Subnet } from "../../types/models";
import { convertNsToHMS, convertHMStoNs } from "../../lib/lib";
import { useNavigate } from "react-router-dom";

export function Form({ subnet }: { subnet?: Subnet }) {
  const [error, setError] = useState<string | null>(null);

  const navigator = useNavigate();

  const scanner_interval = subnet?.scanner_interval
    ? convertNsToHMS(subnet.scanner_interval)
    : undefined;

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = new FormData(e.currentTarget);
    const hours = Number(form.get("scanner_hours") || 0);
    const minutes = Number(form.get("scanner_minutes") || 0);
    const seconds = Number(form.get("scanner_seconds") || 0);

    const durationNs = convertHMStoNs(hours, minutes, seconds);

    const payload = {
      subnet: form.get("subnet"),
      comment: form.get("comment"),
      scanner_enabled: form.get("scanner_enabled") === "true",
      scanner_interval: durationNs,
    };

    const res = subnet
      ? await fetch(`/api/v1/subnet/${subnet.id}`, {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(payload),
        })
      : await fetch("/api/v1/subnet", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(payload),
        });

    if (res.ok) {
      const newSubnet = await res.json();
      navigator(`/subnet/${newSubnet.ID}`);
    } else {
      const errText = await res.text();
      setError(errText);
    }
  };
  return (
    <div className="container">
      {error && <strong style={{ color: "red" }}>{error}</strong>}
      <form onSubmit={handleSubmit}>
        <fieldset>
          <label>
            Subnet
            <input
              type="text"
              name="subnet"
              defaultValue={subnet?.subnet}
              placeholder="192.168.0.0/24"
              required
            />
          </label>
          <label>
            Comment
            <input
              type="text"
              name="comment"
              defaultValue={subnet?.comment}
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
              defaultChecked={subnet ? subnet.scanner_enabled : true}
            />
            Activate
          </label>
          <div className="grid">
            <label>
              Hours
              <input
                type="number"
                name="scanner_hours"
                defaultValue={scanner_interval?.hours || 0}
                min={0}
              />
            </label>
            <label>
              Minutes
              <input
                type="number"
                name="scanner_minutes"
                defaultValue={scanner_interval?.minutes || 5}
                min={0}
                max={59}
              />
            </label>
            <label>
              Seconds
              <input
                type="number"
                name="scanner_seconds"
                defaultValue={scanner_interval?.seconds || 0}
                min={0}
                max={59}
              />
            </label>
          </div>
        </fieldset>
        <button type="submit">{subnet ? "Save" : "Create"}</button>
      </form>
    </div>
  );
}
