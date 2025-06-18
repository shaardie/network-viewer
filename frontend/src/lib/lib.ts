import { format } from "date-fns";
import { de } from "date-fns/locale";

export function formatDurationNs(nanoseconds: number): string {
  const totalSeconds = Math.floor(nanoseconds / 1_000_000_000);
  const hours = Math.floor(totalSeconds / 3600);
  const minutes = Math.floor((totalSeconds % 3600) / 60);
  const seconds = totalSeconds % 60;
  const parts = [];
  if (hours) parts.push(`${hours}h`);
  if (minutes) parts.push(`${minutes}m`);
  parts.push(`${seconds}s`);
  return parts.join(" ");
}

export function formatDateTime(isoString: string): string {
  const date = new Date(isoString);
  return format(date, "Pp", { locale: de });
}

export function convertNsToHMS(ns: number): {
  hours: number;
  minutes: number;
  seconds: number;
} {
  const totalSeconds = Math.floor(ns / 1_000_000_000);
  const hours = Math.floor(totalSeconds / 3600);
  const minutes = Math.floor((totalSeconds % 3600) / 60);
  const seconds = totalSeconds % 60;
  return { hours, minutes, seconds };
}

export function convertHMStoNs(
  hours: number,
  minutes: number,
  seconds: number,
): number {
  return (hours * 3600 + minutes * 60 + seconds) * 1_000_000_000;
}
