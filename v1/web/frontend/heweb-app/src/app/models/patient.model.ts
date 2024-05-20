export interface Message {
  message: string;
  date: string;
}

export interface Medication {
  name: string;
  dosage: number;
  frequency: number;
  isTaken: boolean;
}

export interface Patient {
  _id: string;        // Corresponds to "_id" from MongoDB but is used as 'id' in TypeScript
  name: string;
  email: string;
  phone: string;
  birthday: string;
  gender: string;
  asusvivowatchsn: string;
  photosticker: string;
  messages: Message[];
  medications: Medication[];
}