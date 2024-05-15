export interface Medication {
  Name: string;
  Dosage: number;
  Frequency: number;
}

export interface Patient {
  ID: string;
  Name: string;
  Email: string;
  Phone: string;
  Birthday: string;
  Gender: string;
  PhotoSticker: string;
  Messages: Record<string, string>;
  Medications: Medication[];
}
