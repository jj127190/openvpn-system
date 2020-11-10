$.fn.xcity = function(pName,cName,aName){

    this.p = $(this).find('select[lay-filter=province]');

    this.c = $(this).find('select[lay-filter=city]');

    this.a = $(this).find('select[lay-filter=area]');

    this.cityList = [];

    this.reaList = [];

    this.showP  = function(provinceList) {

        this.p.html('');

        is_pName = false;
        
        for (var i in provinceList) {
            
            if(pName==provinceList[i].name){
                is_pName = true;
                this.cityList = provinceList[i].cityList;
                this.p.append("<option selected value='"+provinceList[i].name+"'>"+provinceList[i].name+"</option>")
            }else{
                this.p.append("<option value='"+provinceList[i].name+"'>"+provinceList[i].name+"</option>")
            }
        }
        if(!is_pName){
            this.cityList = provinceList[0].cityList;
        }
        
    }

    this.showC = function (cityList) {

        this.c.html('');

        is_cName = false;

        for (var i in cityList) {
            if(cName==cityList[i].name){
                is_cName = true;
                this.areaList = cityList[i].areaList;
                this.c.append("<option selected value='"+cityList[i].name+"'>"+cityList[i].name+"</option>")
            }else{
                this.c.append("<option value='"+cityList[i].name+"'>"+cityList[i].name+"</option>")
            }
        }

        if(!is_cName){
            this.areaList = cityList[0].areaList;
        }
    }

    this.showA = function (areaList) {
        this.a.html('');

        for (var i in areaList) {
            
            if(aName==areaList[i]){
                this.a.append("<option selected value='"+areaList[i]+"'>"+areaList[i]+"</option>")
            }else{
                this.a.append("<option value='"+areaList[i]+"'>"+areaList[i]+"</option>")
            }
        }
    }

    this.showP(provinceList);
    this.showC(this.cityList);
    this.showA(this.areaList);

    form.render('select');

    form.on('select(province)', function(data){
        var pName = data.value;
        $(data.elem).parents(".x-city").xcity(pName);
    });

    form.on('select(city)', function(data){
        var cName = data.value;
        var pName = $(data.elem).parents(".x-city").find('select[lay-filter=province]').val();
        console.log(pName);
        $(data.elem).parents(".x-city").xcity(pName,cName);
    });

    return this;
}
var provinceList = [
{name:'背景', cityList:[]},
{name:'上海', cityList:[]},
{name:'天津', cityList:[		   
{name:'市辖区', areaList:['和平区','河东区','河西区','南开区','河北区','红桥区','塘沽区','汉沽区','大港区','东丽区','西青区','津南区','北辰区','武清区','宝坻区']},		   
{name:'县', areaList:['宁河县','静海县','蓟　县']}
]},


];
