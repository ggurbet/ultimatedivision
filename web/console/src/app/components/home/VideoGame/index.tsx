// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { JoinButton } from '@components/common/JoinButton';

import videoPreview from '@static/img/gameLanding/video/video-preview.webp';
import playButton from '@static/img/gameLanding/video/play-button.svg';
import videoBg from '@static/img/gameLanding/video/video-bg.png';

import './index.scss';

/** TODO: add video */

export const VideoGame: React.FC = () =>
    <section className="video-game">
        <div className="video-game__video">
            <div className="video-game__video__preview">
                <img src={videoPreview} alt="video-preview" className="video-game__video__preview__image" />
                <img src={videoBg} alt="video-bg" className="video-game__video__preview__bg" />
                <button className="video-game__video__preview__play-button">
                    <img className="video-game__video__preview__play-button__image" src={playButton} alt="play-button" />
                </button>
            </div>
        </div>
        <JoinButton />
    </section>;
