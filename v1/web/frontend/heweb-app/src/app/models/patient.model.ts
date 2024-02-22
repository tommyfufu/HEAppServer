export interface Patient {
  ID: string;             
  Name: string;           
  Email: string;          
  Phone: string;          
  Birthday: string;       
  PhotoSticker: string;   
  Messages: Record<string, string>; 
  Medication: string[];   
}
