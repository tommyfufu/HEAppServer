export interface Patient {
    id: string;
    name: string;
    phone: string;
    email: string;
    photoSticker: string;
    messages: Record<string, string>;
    medication: string[];
  }
  