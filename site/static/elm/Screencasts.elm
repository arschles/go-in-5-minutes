module Screencasts exposing (..)

import Browser
import Html exposing (Html, text, pre, div)
import Html.Parser
import Html.Parser.Util
import Http



-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL


type Model
  = Failure
  | Loading
  | Success String


init : () -> (Model, Cmd Msg)
init _ =
  ( Loading
  , Http.get
      { url = "/api/v1/screencasts/summary_list"
      , expect = Http.expectString GotText
      }
  )



-- UPDATE


type Msg
  = GotText (Result Http.Error String)


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          (Success fullText, Cmd.none)

        Err _ ->
          (Failure, Cmd.none)



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  case model of
    Failure ->
      text "I was unable to get screencasts."

    Loading ->
      text "Loading..."

    Success rawHtml ->
      let
        nodes =
          case Html.Parser.run rawHtml of
              Ok parsedNodes -> Html.Parser.Util.toVirtualDom parsedNodes
              _ -> 
                [text "Nothing here!"]
      in
      div [] nodes
