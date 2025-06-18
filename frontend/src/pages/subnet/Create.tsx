import { useState } from "react";

export function Create() {
  const [error, setError] = useState<string | null>(null);
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
      <h1>Create Subnet</h1>
      <div className="container">
        {error && <strong style={{ color: "red" }}>{error}</strong>}
        <form onSubmit={handleSubmit}>
          <fieldset>
            <label>
              Subnet
              <input
                type="text"
                name="subnet"
                placeholder="192.168.0.0/24"
                required
              />
            </label>
            <label>
              Comment
              <input type="text" name="comment" placeholder="Comment" />
            </label>
          </fieldset>
          <fieldset>
            <legend>Scanner</legend>
            <label>
              <input
                type="checkbox"
                name="scanner_enabled"
                value="true"
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
          <button type="submit">Create Subnet</button>
        </form>
      </div>
    </>
  );
}
