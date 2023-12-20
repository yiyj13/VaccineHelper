import { useState, useEffect, useMemo } from 'react'
import { Button, Menu } from '@nutui/nutui-react-taro'
import { IconFont, Edit } from '@nutui/icons-react-taro'
import Taro from '@tarojs/taro'

import api, { getProfiles } from '../../api'

import { Profile, TemperatureRecord } from '../../api/methods'
import useSWR from 'swr'

export default function TemperHistory() {
  const [temperRecordList, setTemperRecordList] = useState<TemperatureRecord[]>([])
  const { data: profiles } = useSWR('getProfiles', getProfiles)
  const [menuValue, setMenuValue] = useState<number>(0)

  const fetchTemperatureRecords = async () => {
    const res = await api.get('/api/temperature-records')
    var result = res.data as TemperatureRecord[]
    result.sort((a, b) => new Date(a.date).getTime() - new Date(b.date).getTime())
    setTemperRecordList(result)
  }

  useEffect(() => {
    fetchTemperatureRecords()
  }, [menuValue])

  const handleAddRecord = () => {
    Taro.navigateTo({ url: '/pages/temper/index' })
  }

  const MemberDataList = useMemo(() => {
    const dataList = profiles ? profiles.map((item) => ({ value: item.ID, text: item.relationship })) : []
    dataList.unshift({ value: 0, text: '所有成员' })
    return dataList
  }, [profiles])

  return (
    <div style={{ height: '100%', position: 'relative', paddingBottom: '70px' }}>
      <Menu>
        <Menu.Item
          options={MemberDataList}
          defaultValue={0}
          value={menuValue}
          onChange={(v) => {
            setMenuValue(v.value)
          }}
        />
      </Menu>
      {temperRecordList
        ?.filter((item) => (menuValue !== 0 ? item.profileId === menuValue : true))
        .map((item, index) => (
          <ItemRender data={item} key={index} menuVal={menuValue} />
        ))}
      <Button
        type='primary'
        onClick={handleAddRecord}
        style={{ width: '90%', position: 'fixed', bottom: '10px', marginLeft: '5%', marginRight: '5%', zIndex: 999 }}
      >
        新增记录
      </Button>
    </div>
  )
}

const ItemRender = ({ data, menuVal }: { data: TemperatureRecord; menuVal: number }) => {
  const [profileInfo, setProfileInfo] = useState<Profile>({} as Profile)

  const getProfileInfo = (profileId: number) => {
    api.request({ url: '/api/profiles/' + profileId }).then((res) => {
      const result = res.data as Profile
      setProfileInfo(result)
    })
  }

  useEffect(() => {
    const fetchData = async () => {
      if (data.profileId === menuVal || menuVal === 0) {
        await getProfileInfo(data.profileId)
      }
    }
    fetchData()
  }, [data.profileId, menuVal])

  const handleEditRecord = (recordData: TemperatureRecord) => {
    Taro.navigateTo({
      url: `/pages/temper/index?id=` + recordData.ID,
    })
  }

  return (
    <div
      className='border border-gray-300 p-4 rounded-md'
      style={{ borderRadius: '8px', marginLeft: '10px', marginRight: '10px' }}
    >
      <div className='flex items-center justify-between'>
        <div className='flex items-center'>
          <IconFont className='text-2xl mr-2' name={profileInfo.avatar} style={{ width: '40px', height: '40px' }} />
          <div className='flex justify-between mt-2'>
            <div className='font-bold' style={{ color: '#4796A1' }}>
              {profileInfo.relationship}
            </div>
            <div className='font-bold ml-2'>{profileInfo.fullName}</div>
          </div>
        </div>
        <Edit className='cursor-pointer' onClick={() => handleEditRecord(data)} />
      </div>
      <div className='flex justify-between mt-2'>
        <div className='text-gray-500'>
          测温时间 <b className='text-black font-bold'>{data.date}</b>
        </div>
        <div className='text-gray-500'>
          体温值 <b className='text-black font-bold'>{data.temperature.toFixed(1)}</b>
        </div>
      </div>
    </div>
  )
}