export type Subnet = {
  id: number;
  subnet: string;
  scanner_enabled: boolean;
  scanner_interval: number;
  last_scan: string;
  comment: string;
};

export type IP = {
  id: number;
  ip: string;
  hostname: string;
  online: boolean;
  rtt: number;
  comment: string;
};
