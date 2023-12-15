/* TODO: 
    1. Delay in data update after a picker is checked
    2. CSS style for the menu and the buttons
*/

import Taro from '@tarojs/taro'
import { useState, useEffect } from 'react'
import { Cell, Switch, Picker, Uploader, Button, DatePicker, TextArea, Input } from '@nutui/nutui-react-taro'
import { PickerOption } from '@nutui/nutui-react-taro/dist/types/packages/picker/types'

import ComboBox from '../../components/combobox'
import api from '../../api'
import { Vaccine, Profile, RecordData } from '../../api/methods'

export default function VaccineRecord() {
  const [record, setRecord] = useState<Partial<RecordData>>({
    vaccinationDate: '',
    type: '',
    reminder: false,
    remindDate: '',
    nextVaccinationDate: '',
    voucher: '',
    note: '',
  })

  const [VaccineData, setVaccineData] = useState<PickerOption[]>([])
  const [MemberData, setMemberData] = useState<PickerOption[]>([])

  useEffect(() => {
    api.get('/api/vaccines').then((res) => {
      let vacData = res.data as Vaccine[]
      setVaccineData(vacData.map((item) => ({ value: item.ID, text: item.name })))
    })

    api.get('/api/profiles').then((res) => {
      let memData = res.data as Profile[]
      setMemberData(memData.map((item) => ({ value: item.ID, text: item.relationship })))
    })
  }, [record])

  const TypeData = [
    [
      { value: 0, text: '免疫接种第一针' },
      { value: 1, text: '免疫接种第二针' },
      { value: 2, text: '免疫接种第三针' },
      { value: 3, text: '免疫接种第四针' },
      { value: 4, text: '常规接种' },
      { value: 5, text: '加强针剂' },
      { value: 6, text: '补种疫苗' },
    ],
  ]

  const ValidData = [
    [
      { value: 0, text: '3月' },
      { value: 1, text: '6月' },
      { value: 2, text: '9月' },
      { value: 3, text: '1年' },
      { value: 4, text: '2年' },
      { value: 5, text: '3年' },
      { value: 6, text: '4年' },
      { value: 7, text: '5年' },
      { value: 8, text: '6年' },
      { value: 9, text: '7年' },
      { value: 10, text: '8年' },
      { value: 11, text: '9年' },
      { value: 12, text: '10年' },
      { value: 13, text: '终身' },
    ],
  ]

  const [idVisible, setIdVisible] = useState(false)
  const [idDesc, setIdDesc] = useState('')
  const confirmID = (options: PickerOption[], values: (string | number)[]) => {
    let description = ''
    options.forEach((option: any) => {
      description += option.text
    })
    setIdDesc(description)
    setRecord({
      ...record,
      profileId: values[0] as number,
    })
  }

  const [nameVisible, setNameVisible] = useState(false)
  const [nameDesc, setNameDesc] = useState('')
  const confirmName = (options: PickerOption[], values: (string | number)[]) => {
    let description = ''
    options.forEach((option: any) => {
      description += option.text
    })
    setNameDesc(description)
    setRecord({
      ...record,
      vaccineId: values[0] as number,
    })
  }

  const [typeVisible, setTypeVisible] = useState(false)
  const [typeDesc, setTypeDesc] = useState('')
  const confirmType = (options: PickerOption[]) => {
    let description = ''
    options.forEach((option: any) => {
      description += option.text
    })
    setTypeDesc(description)
    setRecord({
      ...record,
      type: description,
    })
  }

  const startDate = new Date(2000, 0, 1)
  const endDate = new Date(2025, 11, 30)
  const [dateVisible, setDateVisible] = useState(false)
  const [dateDesc, setDateDesc] = useState('')
  const confirmDate = (values: (string | number)[], _options: PickerOption[]) => {
    const date = values.slice(0, 3).join('-')
    setDateDesc(`${date}`)
    setRecord({
      ...record,
      vaccinationDate: date,
    })
    console.log(record.vaccinationDate)
  }

  const [validVisible, setValidVisible] = useState(false)
  const [validDesc, setValidDesc] = useState('')

  const calculateNextVacDate = () => {
    if (record.vaccinationDate !== undefined) {
      const vacDate = record.vaccinationDate
      console.log('vacDate:', vacDate)
      console.log('validDesc:', validDesc)
      const nextVacDate = addDays(vacDate, validDesc)
      console.log('nextVacDate:', nextVacDate)
      setRecord({
        ...record,
        nextVaccinationDate: nextVacDate,
      })
    }
  }

  const addDays = (dateString: string, days: string) => {
    const dateArray = dateString.split('-')
    const year = parseInt(dateArray[0], 10)
    const month = parseInt(dateArray[1], 10)
    const day = parseInt(dateArray[2], 10)
    if (days.slice(-1) === '月') {
      const addMonth = parseInt(days.slice(0, -1), 10)
      const newMonth = month + addMonth
      return `${year}-${newMonth}-${day}`
    } else if (days.slice(-1) === '年') {
      const addYear = parseInt(days.slice(0, -1), 10)
      const newYear = year + addYear
      return `${newYear}-${month}-${day}`
    } else if (days.slice(-1) === '终') {
      return '终身有效'
    } else {
      return 'Error' // for debug
    }
  }

  const confirmValid = (options: PickerOption[], _values: (string | number)[]) => {
    let description = ''
    options.forEach((option: any) => {
      description += option.text
    })
    setValidDesc(description)
    console.log('Valid description:', description)
    console.log('valid desc:', validDesc)
    calculateNextVacDate()
  }

  const [remindSwitch, setRemindSwitch] = useState(false)
  const [remindVisible, setRemindVisible] = useState(false)
  const [remindValue, setRemindValue] = useState('')
  const [remindUnit, setRemindUnit] = useState('')

  const onSwitchChange = (value: boolean) => {
    setRemindSwitch(value)
    setRecord({
      ...record,
      reminder: value,
    })
    setRemindVisible(value)
  }

  const onRemindValueChange = (value: string) => {
    setRemindValue(value)
  }

  const onRemindUnitSet = (option: string) => {
    setRemindUnit(option)

    calculateRemindDate()
  }

  const [noteValue, setNoteValue] = useState('')
  const onNoteChange = (value: string) => {
    setRecord({
      ...record,
      note: value,
    })
    setNoteValue(value)
  }

  const subtractDays = (dateString: string, days: string) => {
    const dateArray = dateString.split('-')
    const year = parseInt(dateArray[0], 10)
    const month = parseInt(dateArray[1], 10)
    const day = parseInt(dateArray[2], 10)

    if (days.slice(-1) === '日') {
      const subtractDay = parseInt(days.slice(0, -1), 10)
      const newDay = day - subtractDay

      if (newDay > 0) {
        return `${year}-${month}-${newDay}`
      } else {
        const newMonth = month - 1
        const daysInPreviousMonth = new Date(year, newMonth, 0).getDate()
        const correctedDay = daysInPreviousMonth + newDay
        return `${year}-${newMonth}-${correctedDay}`
      }
    } else if (days.slice(-1) === '周') {
      const subtractWeek = parseInt(days.slice(0, -1), 10)
      const newDay = day - subtractWeek * 7

      if (newDay > 0) {
        return `${year}-${month}-${newDay}`
      } else {
        const newMonth = month - 1
        const daysInPreviousMonth = new Date(year, newMonth, 0).getDate()
        const correctedDay = daysInPreviousMonth + newDay
        return `${year}-${newMonth}-${correctedDay}`
      }
    } else if (days.slice(-1) === '月') {
      const subtractMonth = parseInt(days.slice(0, -1), 10)
      const newMonth = month - subtractMonth

      if (newMonth > 0) {
        return `${year}-${newMonth}-${day}`
      } else {
        const newYear = year - 1
        const correctedMonth = 12 + newMonth
        return `${newYear}-${correctedMonth}-${day}`
      }
    } else {
      return 'Error' // for debug
    }
  }

  const calculateRemindDate = () => {
    if (
      remindValue !== undefined &&
      !isNaN(parseInt(remindValue)) &&
      remindUnit !== undefined &&
      remindUnit !== '' &&
      record.nextVaccinationDate !== undefined &&
      record.remindDate !== undefined
    ) {
      const remindDate = subtractDays(record.nextVaccinationDate, remindValue + remindUnit)
      setRecord({
        ...record,
        remindDate: remindDate,
      })
    } else {
      setRecord({
        ...record,
        remindDate: '',
      })
    }
  }

  const handleSubmission = async () => {
    console.log(record) // for debug
    if (
      record &&
      record.profileId !== undefined &&
      record.profileId >= 0 &&
      record.vaccineId !== undefined &&
      record.vaccineId >= 0 &&
      record.type !== undefined &&
      record.vaccinationDate
    ) {
      try {
        const res = await api.request({ url: '/api/vaccination-records', method: 'POST', data: record })
        console.log(res.data) // for debug
        Taro.showToast({ title: '提交成功', icon: 'success' })
        setTimeout(() => {
          Taro.navigateBack()
        }, 1000)
      } catch (error) {
        console.log('Error submitting vaccination record:', error)
        Taro.showToast({ title: '提交失败', icon: 'error' })
      }
    } else {
      Taro.showToast({ title: '请填写完整记录', icon: 'error' })
    }
  }

  const handleReset = () => {
    setRecord({
      vaccinationDate: '',
      type: '',
      reminder: false,
      remindDate: '',
      nextVaccinationDate: '',
      voucher: '',
      note: '',
    })

    setIdDesc('')
    setNameDesc('')
    setTypeDesc('')
    setDateDesc('')
    setValidDesc('')
    setRemindValue('')
    setRemindUnit('')
    setNoteValue('')
    setRemindSwitch(false)
  
    setIdVisible(false)
    setNameVisible(false)
    setTypeVisible(false)
    setDateVisible(false)
    setValidVisible(false)
    setRemindVisible(false)

    Taro.showToast({ title: '重置成功', icon: 'success' })
  }

  return (
    <>
      <Cell
        title='接种人'
        description={idDesc}
        onClick={() => setIdVisible(!idVisible)}
        style={{ textAlign: 'center' }}
      />
      <Picker
        title='接种人'
        visible={idVisible}
        options={MemberData}
        onConfirm={(list, values) => confirmID(list, values)}
        onClose={() => setIdVisible(false)}
      />
      <Cell
        title='疫苗名称'
        description={nameDesc}
        onClick={() => setNameVisible(!nameVisible)}
        style={{ textAlign: 'center' }}
      />
      <Picker
        title='疫苗名称'
        visible={nameVisible}
        options={VaccineData}
        onConfirm={(list, values) => confirmName(list, values)}
        onClose={() => setNameVisible(false)}
      />
      <Cell
        title='接种类型'
        description={typeDesc}
        onClick={() => setTypeVisible(!typeVisible)}
        style={{ textAlign: 'center' }}
      />
      <Picker
        title='接种类型'
        visible={typeVisible}
        options={TypeData}
        onConfirm={(list) => confirmType(list)}
        onClose={() => setTypeVisible(false)}
      />
      <Cell
        title='接种时间'
        description={dateDesc}
        onClick={() => setDateVisible(true)}
        style={{ textAlign: 'center' }}
      />
      <DatePicker
        title='接种时间'
        startDate={startDate}
        endDate={endDate}
        visible={dateVisible}
        type='date'
        onClose={() => setDateVisible(false)}
        onConfirm={(options, values) => confirmDate(values, options)}
      />
      <Cell
        title='有效期限'
        description={validDesc}
        onClick={() => setValidVisible(!validVisible)}
        style={{ textAlign: 'center' }}
      />
      <Picker
        title='有效期限'
        visible={validVisible}
        options={ValidData}
        onConfirm={(list, values) => confirmValid(list, values)}
        onClose={() => setValidVisible(false)}
      />
      <div className='col-span-full flex-content flex items-center '>
        <span className='text-sm ml-2 '>接种提醒</span>
        <Switch className=' ml-2' checked={remindSwitch} onChange={(value) => onSwitchChange(value)} />
        {remindVisible && (
          <div className='ml-2 flex items-center'>
            <Input
              type='number'
              placeholder='数字'
              value={remindValue}
              onChange={(value) => onRemindValueChange(value)}
              className='mr-2 border border-gray-200 rounded p-1'
            />
            <ComboBox
              title={''}
              options={['日', '周', '月']}
              onSelect={(option) => onRemindUnitSet(option)}
              className='mr-2'
            />
            <span className='text-sm ml-2'>前提醒</span>
          </div>
        )}
      </div>

      <div className='col-span-full flex-content items-center'>
        <span className='ml-2 text-sm'>接种凭证</span>
        <Uploader
          className='w-full px-2'
          url='https://img13.360buyimg.com/imagetools/jfs/t1/169186/5/33010/1762/639703a1E898bcb96/6c372c661c6dddfe.png'
        />
      </div>
      <Cell title='TextArea' className='col-span-full px-8' style={{ borderRadius: '8px' }}>
        <TextArea placeholder='请输入备注' value={noteValue} autoSize onChange={(value) => onNoteChange(value)} />
      </Cell>
      <div className='col-span-full flex justify-center mt-4'>
        <Button className='submit_btm' formType='submit' type='primary' onClick={handleSubmission}>
          提交
        </Button>
        <div style={{ marginLeft: '16px' }}>
          <Button className='reset_btm' formType='reset' onClick={handleReset}>
            重置
          </Button>
        </div>
      </div>
    </>
  )
}