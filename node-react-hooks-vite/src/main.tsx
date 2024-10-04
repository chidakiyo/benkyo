import { createContext, render } from 'preact'
import { App } from './app.tsx'
import './index.css'
import React from 'preact/compat'
import PersonalInfoContext, { personalInfo } from './info.tsx';

render(
    <PersonalInfoContext.Provider value={personalInfo}>
        <React.StrictMode>
            <App />
        </React.StrictMode>
    </PersonalInfoContext.Provider>
    , document.getElementById('app')!
);


    