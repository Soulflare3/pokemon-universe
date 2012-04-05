/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

import (
	"time"
	"math/rand"
	
	"putools/log"
)

type TimeService struct {
	isRunning			bool
	lastWeatherUpdate	int64
	forcedWeather		int
	
	Day			int
	Hour		int
	Minutes		int
	
	Weather		int
}

func NewTimeService() *TimeService {
	timeService := &TimeService { isRunning: false,
								  lastWeatherUpdate: PUSYS_TIME(),
								  forcedWeather: 9,
								  Day: 0,
								  Hour: 0,
								  Minutes: 0,
								  Weather: 0 }
	
	return timeService
}

func (t *TimeService) GenerateWeather() {
	weather := t.forcedWeather
	if weather == 9 {
		weather = rand.Intn(10)
	}
	
	switch weather {
		case 0:
			t.Weather = WEATHER_NORMAL
		case 1:
			t.Weather = WEATHER_RAIN
		case 2:
			t.Weather = WEATHER_HAIL
		case 3:
			t.Weather = WEATHER_FOG
		case 4:
			t.Weather = WEATHER_SANDSTORM
		default:
			t.Weather = WEATHER_NORMAL
	}
	
	logger.Printf("INFO: Weather type changed to: %v\n", t.WeatherToStr())
}

func (t *TimeService) Start() {
	if !t.isRunning {
		// Calculate in-game time from real world time
		// this way we always have the right in-game time
		t.calculateTimeFromSystem()
	
		// Generate random starting weather
		t.GenerateWeather()
		
		t.isRunning = true
		
		t.run()
		logger.Println("INFO: TimeService started")
	}
}

func (t *TimeService) Stop() {
	if t.isRunning {	
		logger.Println("INFO: TimeServie stopping")
		t.isRunning = false
	}
}

func (t *TimeService) IsNight() bool {
	return (t.Hour >= 20 || t.Hour < 6)
}

func (t *TimeService) SetForcedWeather(_weather int) {
	t.forcedWeather = _weather
	t.lastWeatherUpdate = 0
}

func (t *TimeService) WeatherToStr() (str string) {
	str = "Normal"
	switch t.Weather {
		case 1:
			str = "Rain"
		case 2:
			str = "Hail"
		case 3:
			str = "Fog"
		case 4:
			str = "Sandstorm"
	}
	
	return
}

func (t *TimeService) calculateTimeFromSystem() {
	currentTime := time.Now()
	currentDay := currentTime.Weekday()
	currentHour := currentTime.Hour()
	currentMinute := currentTime.Minute()
	
	if currentDay == time.Monday {
		if currentHour < 12 {
			t.Day = int(time.Monday)
		} else {
			t.Day = int(time.Tuesday)
		}
	} else if currentDay == time.Tuesday {
		if currentHour < 12 {
			t.Day = int(time.Wednesday)
		} else {
			t.Day = int(time.Thursday)
		}
	} else if currentDay == time.Wednesday {
		if currentHour < 12 {
			t.Day = int(time.Friday)
		} else {
			t.Day = int(time.Saturday)
		}
	} else if currentDay == time.Thursday {
		if currentHour < 12 {
			t.Day = int(time.Sunday)
		} else {
			t.Day = int(time.Monday)
		}
	} else if currentDay == time.Friday {
		if currentHour < 12 {
			t.Day = int(time.Tuesday)
		} else {
			t.Day = int(time.Wednesday)
		}
	} else if currentDay == time.Saturday {
		if currentHour < 12 {
			t.Day = int(time.Thursday)
		} else {
			t.Day = int(time.Friday)
		}
	} else {
		if currentHour < 12 {
			t.Day = int(time.Saturday)
		} else {
			t.Day = int(time.Sunday)
		}
	}
	
	if currentHour >= 12 {
		currentHour = currentHour - 12
	}
	t.Hour = currentHour * 2
	
	if currentMinute >= 30 {
		currentMinute = currentMinute - 30
		
		// 30 min real world time, increment in-game hour by 1
		t.Hour++
	}
	t.Minutes = currentMinute * 2
	
	logger.Printf("INFO: Current time is %v @ %d:%d\n", Days[t.Day], t.Hour, t.Minutes)
}

func (t *TimeService) run() {
	if t.isRunning {
		
		// In-game clock runs twice as fast as real time
		go func() {
			time.Sleep(time.Second * 30)
			t.run()
		}()
		
		currentMin := t.Minutes
		if currentMin == 59 {
			currentMin = 0
			currentHour := t.Hour
	
			if currentHour == 23 {
				currentHour = 0
				
				t.incrementDay()
			} else {
				currentHour++
			}
			
			t.Hour = currentHour
		} else {
			currentMin++
		}
		t.Minutes = currentMin
		
		weatherUpdate := PUSYS_TIME() - t.lastWeatherUpdate
		if weatherUpdate > 3600000 {
			t.GenerateWeather()
			t.lastWeatherUpdate = PUSYS_TIME()
		}

	} else {
		logger.Println("INFO: TimeServie stopped")
	}
}

func (t *TimeService) incrementDay() {
	currentDay := t.Day
	
	if currentDay == 6 {
		currentDay = 0
	} else {
		currentDay++
	}
	
	t.Day = currentDay
}