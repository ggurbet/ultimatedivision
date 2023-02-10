// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import Unity, { UnityContent } from 'react-unity-webgl';

import './index.scss';

const FootballGame: React.FC = () => {
    const unityContext = new UnityContent('/static/dist/webGl/football.json', '/static/dist/webGl/UnityLoader.js');

    return <Unity unityContent={unityContext} className="unity-container" />;
};

export default FootballGame;
