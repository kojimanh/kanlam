package ping_use_case

import (
	pingdto "kanlam/entity/dto/ping.dto"
	storageping "kanlam/storage/maindb/ping.storage"
)

type PingUseCase struct {
}

func (r PingUseCase) Ping() (pingdto.PingOutputDto, error) {
	greeting, err := storageping.PingDb()

	result := pingdto.PingOutputDto{
		Result: greeting,
	}

	return result, err
}
