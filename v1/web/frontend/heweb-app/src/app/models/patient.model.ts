export interface Patient {
  ID: string;             // Matches "ID" from the JSON
  Name: string;           // Matches "Name" from the JSON
  Email: string;          // Matches "Email" from the JSON
  Phone: string;          // Matches "Phone" from the JSON
  Birthday: string;       // Add this property to match the JSON
  PhotoSticker: string;   // Matches "PhotoSticker" from the JSON
  Messages: Record<string, string>;  // Matches "Messages" from the JSON
  Medication: string[];   // Matches "Medication" from the JSON
}
