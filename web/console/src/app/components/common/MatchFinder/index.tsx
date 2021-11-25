// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect, useMemo } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { toast } from 'react-toastify';

import { Timer } from './Timer';
import { AutoCloseTimer } from './AutoCloseTimer';

import { QueueClient } from '@/api/queue';
import { RouteConfig } from '@/app/routes';
import { RootState } from '@/app/store';
import { getMatchScore } from '@/app/store/actions/mathes';
import { startSearchingMatch } from '@/app/store/actions/clubs';

import './index.scss';

const MatchFinder: React.FC = () => {
    const { squad } = useSelector((state: RootState) => state.clubsReducer.activeClub);
    const { isSearchingMatch } = useSelector((state: RootState) => state.clubsReducer);

    const queueClient = useMemo(() => new QueueClient(), [isSearchingMatch]);

    const dispatch = useDispatch();
    const history = useHistory();

    /** Indicates if match is found. */
    const [isMatchFound, setIsMatchFound] = useState<boolean>(false);

    /** variable describes that user added to gueue */
    const YOU_ADDED_MESSAGE: string = 'you added!';
    /** variable describes that it needs confirm game from user */
    const YOU_CONFIRM_PLAY_MESSAGE: string = 'you confirm play?';
    /** variable describes that was send wrong action from user */
    const WRONG_ACTION_MESSAGE: string = 'wrong action';
    /** variable describes that webscoket connection responsed with error */
    const ERROR_MESSAGE: string = 'could not write to websocket';
    /** variable describes that user still searching game */
    const STILL_SARCHING_MESSAGE: string = 'you are still in search!';

    /** Delay is time delay for redirect user to match page */
    const DELAY: number = 2000;

    /** variables describes first and second teams indexes for eventAction response. */
    const FIRST_TEAM_INDEX: number = 0;
    const SECOND_TEAM_INDEX: number = 1;

    /** canceles searching game and closes MatchFinder component. */
    const canselSearchingGame = () => {
        queueClient.ws.send(JSON.stringify({ action: 'finishSearch' }));
        dispatch(startSearchingMatch(false));
        setIsMatchFound(false);
    };

    /** TODO: it not uses now. Will be reworked after back-end solutions */
    /** canceles confirmation game. */
    const cancelConfirmationGame = () => {
        queueClient.ws.send(JSON.stringify({ action: 'reject' }));
        setIsMatchFound(false);
    };

    /** exposes confirm match logic */
    const confirmMatch = () => {
        queueClient.ws.send(JSON.stringify({ action: 'confirm', squadId: squad.id }));
    };

    /** processes queue client event message */
    queueClient.ws.onmessage = ({ data }) => {
        const eventAction = JSON.parse(data);
        switch (eventAction.message) {
        case YOU_CONFIRM_PLAY_MESSAGE:
            setIsMatchFound(true);

            return;
        case STILL_SARCHING_MESSAGE:
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
        case ERROR_MESSAGE:
            toast.error('Something wrong, please, try later.', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });

            return;
        case YOU_ADDED_MESSAGE:
            setIsMatchFound(false);

            return;
        default:
            const firstTeamGoalsCrored =
                    eventAction.message[FIRST_TEAM_INDEX].quantityGoals;
            const secondTeamGoalsScored =
                    eventAction.message[SECOND_TEAM_INDEX].quantityGoals;

            toast.success('Successfully! You will be redirected to match page', {
                position: toast.POSITION.TOP_RIGHT,
            });

            dispatch(getMatchScore({ firstTeamGoalsCrored, secondTeamGoalsScored }));
            dispatch(startSearchingMatch(false));

            setTimeout(() => {
                history.push(RouteConfig.Match.path);
            }, DELAY);
        }
    };

    /** exposes start searching match logic. */
    const startSearchMatch = () => {
        queueClient.startSearch('startSearch', squad.id);
    };

    queueClient.ws.onerror = (event: Event) => {
        toast.error('Something wrong, please, try later.', {
            position: toast.POSITION.TOP_RIGHT,
            theme: 'colored',
        });
    };

    useEffect(() => {
        isSearchingMatch && startSearchMatch();
    }, [isSearchingMatch]);

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
                    value="Not Allow"
                    type="button"
                />}
                {isMatchFound ? <input
                    className="match-finder__form__cancel"
                    value="not allow"
                    type="button"
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
