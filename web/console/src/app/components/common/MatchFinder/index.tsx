// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { toast } from 'react-toastify';

import { AutoCloseTimer } from './AutoCloseTimer';
import { Timer } from './Timer';

import { QueueClient } from '@/api/queue';
import { RouteConfig } from '@/app/routes';
import { RootState } from '@/app/store';
import { getMatchScore } from '@/app/store/actions/mathes';
import { startSearchingMatch } from '@/app/store/actions/clubs';

import './index.scss';

const MatchFinder: React.FC = () => {
    const { squad } = useSelector((state: RootState) => state.clubsReducer.activeClub);
    const { isSearchingMatch } = useSelector((state: RootState) => state.clubsReducer);

    const [queueClient, setQueueClient] = useState<QueueClient>(new QueueClient());

    const dispatch = useDispatch();
    const history = useHistory();

    /** Indicates if match is found. */
    const [isMatchFound, setIsMatchFound] = useState<boolean>(false);

    /** Delay is time delay for redirect user to match page. */
    const DELAY: number = 2000;

    /** variables describes first and second teams indexes for eventAction response. */
    const FIRST_TEAM_INDEX: number = 0;
    const SECOND_TEAM_INDEX: number = 1;

    /** variable describes that webscoket connection responsed with error. */
    const ERROR_MESSAGE: string = 'could not write to websocket';
    /** variable describes that user still searching game. */
    const STILL_SEARCHING_MESSAGE: string = 'you are still in search!';
    /** variable describes that was send wrong action from user. */
    const WRONG_ACTION_MESSAGE: string = 'wrong action';
    /** variable describes that user added to gueue. */
    const YOU_ADDED_MESSAGE: string = 'you added!';
    /** variable describes that it needs confirm game from user. */
    const YOU_CONFIRM_PLAY_MESSAGE: string = 'you confirm play?';
    /** Variable describes that user have leaved from searching game. */
    const YOU_LEAVED_MESSAGE: string = 'you leaved!';

    /** exposes confirm match logic. */
    const confirmMatch = () => {
        queueClient.sendAction('confirm', squad.id);
    };

    /** canceles confirmation game. */
    const cancelConfirmationGame = () => {
        queueClient.sendAction('reject', squad.id);
    };

    /** canceles searching game and closes MatchFinder component. */
    const canselSearchingGame = () => {
        /** TODO: rework after ./queue/chore.go solution. */
        queueClient.ws.close();

        const newQueueClient = new QueueClient();
        newQueueClient.finishSearch('finishSearch', squad.id);

        setQueueClient(newQueueClient);
        dispatch(startSearchingMatch(false));
    };

    /** exposes start searching match logic. */
    const startSearchMatch = () => {
        /** TODO: rework after ./queue/chore.go solution. */
        const newQueueClient = new QueueClient();
        newQueueClient.startSearch('startSearch', squad.id);
        setQueueClient(newQueueClient);
    };

    useEffect(() => {
        isSearchingMatch && startSearchMatch();
    }, [isSearchingMatch]);

    /** processes queue client event messages. */
    queueClient.ws.onmessage = ({ data }: MessageEvent) => {
        const messageEvent = JSON.parse(data);

        switch (messageEvent.message) {
        case ERROR_MESSAGE:
            toast.error('error message', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });

            return;
        case STILL_SEARCHING_MESSAGE:
            /** TODO: will be deleted after ./queue/chore.go reworks. */
            queueClient.ws.close();

            const newQueueClient = new QueueClient();
            newQueueClient.startSearch('startSearch', squad.id);

            setQueueClient(newQueueClient);

            setIsMatchFound(false);

            toast.error('Your game was canceled. You are still in search.', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });

            return;
        case WRONG_ACTION_MESSAGE:
            toast.error('Something wrong, please, try later.', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });

            return;
        case YOU_ADDED_MESSAGE:
            setIsMatchFound(false);

            return;
        case YOU_CONFIRM_PLAY_MESSAGE:
            setIsMatchFound(true);

            return;
        case YOU_LEAVED_MESSAGE:
            dispatch(startSearchingMatch(false));

            return;
        default:
            const firstTeamGoalsCrored =
                    messageEvent.message[FIRST_TEAM_INDEX].quantityGoals;
            const secondTeamGoalsScored =
                    messageEvent.message[SECOND_TEAM_INDEX].quantityGoals;

            toast.success('Successfully! You will be redirected to match page', {
                position: toast.POSITION.TOP_RIGHT,
            });

            dispatch(getMatchScore({ firstTeamGoalsCrored, secondTeamGoalsScored }));
            dispatch(startSearchingMatch(false));

            /** implements redirect to match page after DELAY time.  */
            setTimeout(() => {
                history.push(RouteConfig.Match.path);
            }, DELAY);
        }
    };

    queueClient.ws.onerror = (event: Event) => {
        toast.error('Something wrong, please, try later.', {
            position: toast.POSITION.TOP_RIGHT,
            theme: 'colored',
        });
    };

    return isSearchingMatch && <section className={isMatchFound ? 'match-finder__wrapper' : ''}>
        <div className="match-finder">
            <h1 className="match-finder__title">
                {isMatchFound ? 'YOUR MATCH WAS FOUND' : 'LOOKING FOR A MATCH'}
            </h1>
            {isMatchFound ? <AutoCloseTimer /> : <Timer />}
            <div className="match-finder__form">
                {isMatchFound ? <input
                    className="match-finder__form__accept"
                    value="Accept"
                    type="button"
                    onClick={confirmMatch}
                /> : <input
                    className="match-finder__form__accept-not-allowed"
                    value="Accept"
                    type="button"
                />}
                {isMatchFound ? <input
                    className="match-finder__form__cancel"
                    value="Cancel"
                    type="button"
                    onClick={cancelConfirmationGame}
                /> : <input
                    className="match-finder__form__cancel"
                    value="Cancel"
                    type="button"
                    onClick={canselSearchingGame}
                />}
            </div>
        </div>
    </section>;
};

export default MatchFinder;
