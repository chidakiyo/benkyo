import { createContext } from "preact";

// useContext用のデータ
export const personalInfo = {
    name: "hogetaro",
    age: 100,
};

const PersonalInfoContext = createContext(personalInfo);
export default PersonalInfoContext;
